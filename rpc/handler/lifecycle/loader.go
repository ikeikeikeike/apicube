package lifecycle

import (
	"context"
	"net/url"

	"github.com/facebookgo/inject"

	"github.com/ikeikeikeike/apicube/base/util"
	"github.com/ikeikeikeike/apicube/base/util/logger"
	"github.com/ikeikeikeike/apicube/base/util/rpc"
	pb "github.com/ikeikeikeike/apicube/rpc/protocol/apicube/lifecycle"
)

// Inject injects dependencies
func Inject(ctx context.Context, env util.Environment, g *inject.Graph, rt *rpc.Mux) {
	var ping pingService

	// inject
	err := g.Provide(
		&inject.Object{Value: &ping},
	)
	if err != nil {
		logger.Panicf("[PANIC] failed to process injection: %s", err)
	}
	uri, err := url.Parse(env.EnvString("URI"))
	if err != nil {
		logger.Panicf("failed to get parse uri: %s", err)
	}

	// routes
	pb.RegisterPingServiceServer(rt.GrpcMux, &ping)

	if err := pb.RegisterPingServiceHandlerFromEndpoint(ctx, rt.GwMux, uri.Host, rt.GwOpts); err != nil {
		logger.Panicf("[PANIC] failed to register grpc gateway: %s", err)
	}
}
