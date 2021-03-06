package main

import (
	"context"
	"database/sql"
	"os"
	"runtime"
	"time"

	// use MySQL
	_ "github.com/go-sql-driver/mysql"

	"github.com/facebookgo/inject"
	"github.com/gigawattio/metaflector"
	"github.com/kelseyhightower/envconfig"
	"github.com/olivere/elastic"
	"github.com/spf13/cast"

	dataes "github.com/ikeikeikeike/apicube/base/data/es"
	"github.com/ikeikeikeike/apicube/base/data/rdb"
	"github.com/ikeikeikeike/apicube/base/data/repo"
	"github.com/ikeikeikeike/apicube/base/domain/trans"
	"github.com/ikeikeikeike/apicube/base/domain/usecase"
	"github.com/ikeikeikeike/apicube/base/util"
	"github.com/ikeikeikeike/apicube/base/util/logger"
)

// Env is localenv which receives variable from environment variable.
type (
	Env struct {
		util.Env
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

// Commander is ...
type commander struct {
	Env *util.Env `inject:""`
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // noneed
	logger.SetDebug(true)
	os.Exit(_main(os.Args[1:]))
}

func _main(args []string) int {
	var env util.Environment = &Env{}
	if err := envconfig.Process("", env); err != nil {
		logger.Panicf("failed to process env var: %s", err)
	}

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

	// Run command
	//
	cmd := &commander{}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Minute)
	defer cancel()

	// Inject
	initInject(ctx, env, db, es, cmd)

	return 0
}

func initInject(ctx context.Context, env util.Environment, db *sql.DB, esc *elastic.Client, cmd *commander) {
	// Injects dependecies
	var g inject.Graph
	var rt interface{}

	err := g.Provide(
		&inject.Object{Value: env},
		&inject.Object{Value: db},
		&inject.Object{Value: esc},
	)
	if err != nil {
		logger.Panicf("failed to process inject.Provide: %s", err)
	}

	dataes.Inject(ctx, env, &g, rt)
	rdb.Inject(ctx, env, &g, rt)
	repo.Inject(ctx, env, &g, rt)
	usecase.Inject(ctx, env, &g, rt)
	trans.Inject(ctx, env, &g, rt)

	// inject to commander
	if err := g.Provide(&inject.Object{Value: cmd}); err != nil {
		logger.Panicf("failed to process inject.Provide: %s", err)
	}
	if err := g.Populate(); err != nil {
		logger.Panicf("failed to process inject.Populate: %s", err)
	}
}
