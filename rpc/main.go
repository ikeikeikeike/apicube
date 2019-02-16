package main

import (
	"context"
	"database/sql"
	"net"
	"net/http"
	// use pprof
	_ "net/http/pprof"
	"net/url"
	"os"

	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	channelzsvc "google.golang.org/grpc/channelz/service"
	"google.golang.org/grpc/reflection"

	"github.com/facebookgo/inject"
	"github.com/gigawattio/metaflector"
	"github.com/olivere/elastic"
	// use MySQL
	_ "github.com/go-sql-driver/mysql"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/kelseyhightower/envconfig"
	"github.com/mediocregopher/radix.v2/pool"
	"github.com/spf13/cast"

	"github.com/ikeikeikeike/apicube/base/data/es"
	"github.com/ikeikeikeike/apicube/base/data/rdb"
	"github.com/ikeikeikeike/apicube/base/data/repo"
	"github.com/ikeikeikeike/apicube/base/data/storage"
	"github.com/ikeikeikeike/apicube/base/domain/trans"
	"github.com/ikeikeikeike/apicube/base/domain/usecase"
	"github.com/ikeikeikeike/apicube/base/job"
	"github.com/ikeikeikeike/apicube/base/util"
	utildlm "github.com/ikeikeikeike/apicube/base/util/dlm"
	"github.com/ikeikeikeike/apicube/base/util/graceful"
	"github.com/ikeikeikeike/apicube/base/util/logger"
	"github.com/ikeikeikeike/apicube/base/util/rpc"
	"github.com/ikeikeikeike/apicube/rpc/handler"
)

// Env is localenv which receives variable from environment variable.
type (
	Env struct {
		util.Env

		// GWURI is overwritten. 49152～65535, overritten
		GWURI string `envconfig:"APICUBE_GWURI" default:"http://0.0.0.0:50000"`
		// URI is overwritten. 49152～65535, overritten
		URI string `envconfig:"APICUBE_URI" default:"grpc://0.0.0.0:50001"`

		// Event is consul event's name for deployment.
		VAR string `envconfig:"APICUBE_PRIVATE_VAR" default:"set-something"`
	}
)

// EnvString returns  stirng
func (e *Env) EnvString(prop string) string {
	v, err := cast.ToStringE(metaflector.Get(e, prop))
	if err != nil {
		logger.Cretical("EnvString failed prop `%v`: %s", prop, err)
	}
	return v
}

// EnvInt returns as int
func (e *Env) EnvInt(prop string) int {
	v, err := cast.ToIntE(metaflector.Get(e, prop))
	if err != nil {
		logger.Cretical("EnvInt failed prop `%v`: %s", prop, err)
	}
	return v
}

func main() {
	var env util.Environment = &Env{}
	if err := envconfig.Process("", env); err != nil {
		logger.Panicf("failed to get env var: %s", err)
	}

	logger.SetDebug(env.IsDebug())
	logger.SetSentry(env.IsSentry())

	// Make sure connection established
	//

	// es
	es, err := util.ESConn(env) // defer es.Close()
	if err != nil {
		logger.Panicf("failed to get ESConn: %s", err)
	}
	// db
	db, err := util.DBConn(env)
	if err != nil {
		logger.Panicf("failed to get DBConn: %s", err)
	}
	defer db.Close()
	// redis db
	rdb, err := util.RDBConn(env)
	if err != nil {
		logger.Panicf("failed to get RDBConn: %s", err)
	}
	defer rdb.Empty()
	// DLM
	dlm, err := util.DLMConn(env)
	if err != nil {
		logger.Panicf("failed to get DLMConn: %s", err)
	}
	defer dlm.Close()
	// a server
	uri, err := url.Parse(env.EnvString("URI"))
	if err != nil {
		logger.Panicf("failed to get parse uri: %s", err)
	}
	// a server
	gwURI, err := url.Parse(env.EnvString("GWURI"))
	if err != nil {
		logger.Panicf("failed to get parse gwURI: %s", err)
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// routing
	rt := initRoutes(ctx, env, uri, gwURI)

	// inject
	initInject(ctx, env, rt, db, es, rdb, dlm)

	// Register reflection service on gRPC server.
	if true {
		// Is this for a development?
		reflection.Register(rt.GrpcMux)
		channelzsvc.RegisterChannelzServiceToServer(rt.GrpcMux)
	}

	// Port Listeners
	//

	// profiler
	if env.IsProfiler() {
		go func() {
			logger.Info("profiler started on http://0.0.0.0:6060")
			logger.Println(http.ListenAndServe(":6060", nil))
		}()
	}

	q := make(chan os.Signal)
	signal.Notify(q, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP)
	go func() {
		<-q

		logger.Info("waiting remain hooks to closing...")
		graceful.Shutdown()

		logger.Info("stopping a grpc server...")
		rt.GrpcMux.GracefulStop()
	}()

	errors := make(chan error)

	go func() {
		lis, err := net.Listen("tcp", uri.Host)
		if err != nil {
			logger.Panic("faild to listen: %v", err)
		}

		logger.Info("start grpc server: %s", uri.Host)

		errors <- rt.GrpcMux.Serve(lis)
	}()

	go func() {
		logger.Info("start grpc gateway server: %s", gwURI.Host)

		errors <- http.ListenAndServe(gwURI.Host, rt.GwMux)
	}()

	if err := <-errors; err != nil {
		logger.Panicf("Serve() returned non-nil error on GracefulStop: %v", err)
	}
}

func initRoutes(ctx context.Context, env util.Environment, uri, gwURI *url.URL) *rpc.Mux {
	// opts := []grpc.ServerOption
	// grpc.WriteBufferSize(s int)
	// grpc.ReadBufferSize(s int)
	// grpc.InitialWindowSize(s int32)
	// grpc.InitialConnWindowSize(s int32)
	// grpc.KeepaliveParams(kp keepalive.ServerParameters)
	// grpc.KeepaliveEnforcementPolicy(kep keepalive.EnforcementPolicy)
	// grpc.CustomCodec(codec Codec)
	// grpc.RPCCompressor(cp Compressor)
	// grpc.RPCDecompressor(dc Decompressor)
	// grpc.MaxMsgSize(m int)
	// grpc.MaxRecvMsgSize(m int)
	// grpc.MaxSendMsgSize(m int)
	// grpc.MaxConcurrentStreams(n uint32)
	// grpc.Creds(c credentials.TransportCredentials)
	// grpc.UnaryInterceptor(i UnaryServerInterceptor)
	// grpc.StreamInterceptor(i StreamServerInterceptor)
	// grpc.InTapHandle(h tap.ServerInHandle)
	// grpc.StatsHandler(h stats.Handler)
	// grpc.UnknownServiceHandler(streamHandler StreamHandler)
	// grpc.ConnectionTimeout(d time.Duration)
	// grpc.MaxHeaderListSize(s uint32)

	// XXX: need in the microservice https://github.com/grpc-ecosystem/grpc-opentracing/tree/master/go/otgrpc
	//                               https://github.com/grpc-ecosystem/go-grpc-middleware/tree/master/tracing/opentracing

	// TODO: Prepare to retry behavior which is called by client.
	//
	// opts := []grpc_retry.CallOption{
	// 	grpc_retry.WithBackoff(grpc_retry.BackoffLinear(100 * time.Millisecond)),
	// 	grpc_retry.WithCodes(codes.NotFound, codes.Aborted),
	// }
	// grpc.Dial("myservice.example.com",
	// 	grpc.WithStreamInterceptor(grpc_retry.StreamClientInterceptor()),
	// 	grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor()),
	// )
	//

	// TODO: recovery: https://github.com/grpc-ecosystem/go-grpc-middleware/tree/master/recovery
	// recoveryOpts := []grpc_recovery.Option{
	//     grpc_recovery.WithRecoveryHandler(customFunc),
	// }

	rt := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_validator.StreamServerInterceptor(),
			// grpc_recovery.StreamServerInterceptor(recoveryOpts...),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_validator.UnaryServerInterceptor(),
			// grpc_recovery.UnaryServerInterceptor(recoveryOpts...),
		)),
	)

	gwRT := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		OrigName:     true,
		EmitDefaults: true,
	}))

	return &rpc.Mux{
		GrpcMux: rt,
		GwMux:   gwRT,
		GwOpts:  []grpc.DialOption{grpc.WithInsecure()},
	}
}

func initInject(ctx context.Context, env util.Environment, rt *rpc.Mux, db *sql.DB, esc *elastic.Client, po *pool.Pool, dlm *utildlm.DLM) {
	// Injects dependecies
	var g inject.Graph

	err := g.Provide(
		&inject.Object{Value: env},
		&inject.Object{Value: db},
		&inject.Object{Value: esc},
		&inject.Object{Value: dlm},
		&inject.Object{Value: po},
	)
	if err != nil {
		logger.Panicf("failed to process inject.Provide: %s", err)
	}

	rdb.Inject(ctx, env, &g, rt)
	storage.Inject(ctx, env, &g, rt)
	repo.Inject(ctx, env, &g, rt)
	job.Inject(ctx, env, &g, rt)
	trans.Inject(ctx, env, &g, rt)
	es.Inject(ctx, env, &g, rt)
	usecase.Inject(ctx, env, &g, rt)
	handler.Inject(ctx, env, &g, rt)

	if err := g.Populate(); err != nil {
		logger.Panicf("failed to process inject.Populate: %s", err)
	}
}
