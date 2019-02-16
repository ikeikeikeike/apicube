package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/facebookgo/inject"
	"github.com/gigawattio/metaflector"
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo"
	"github.com/olivere/elastic"
	"github.com/spf13/cast"

	dataes "github.com/ikeikeikeike/apicube/base/data/es"
	"github.com/ikeikeikeike/apicube/base/data/rdb"
	"github.com/ikeikeikeike/apicube/base/data/repo"
	"github.com/ikeikeikeike/apicube/base/data/storage"
	"github.com/ikeikeikeike/apicube/base/util"
	"github.com/ikeikeikeike/apicube/base/util/logger"
)

// Usage
var (
	name = flag.String("idxname", dataes.ProductsName, "index name. products|favorites")
)

func init() {
	flag.Usage = func() {
		msg := "Usage of %s [FLAG] COMMAND\nCOMMAND:\n  reindex\nFLAG: \n"
		fmt.Fprintf(flag.CommandLine.Output(), msg, os.Args[0])
		flag.PrintDefaults()
	}
}

// Env is localenv which receives variable from environment variable.
type (
	Env struct {
		util.Env
		// Event is consul event's name for deployment.
		VAR string `envconfig:"APICUBE_API_PRIVATE_VAR" default:"set-something"`
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
	Env     *util.Env `inject:""`
	Reindex Reindex   `inject:""`
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // noneed
	logger.SetDebug(true)
	os.Exit(_main(os.Args[1:]))
}

func _main(args []string) int {
	flag.Parse()

	if len(args) <= 0 {
		logger.Printf("[ERROR] no args")
		return 1
	}
	index := dataes.Schema(*name)
	if index == "" {
		logger.Printf("[ERROR] A `%s` does not have index content", *name)
		return 2
	}
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
	rc, cmd := &dataes.Result{}, &commander{}

	// Inject
	initInject(env, db, es, cmd)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Minute)
	defer cancel()

	switch strings.Join(args[len(args)-1:], " ") {
	// case "upsert":
	// logger.Printf("[INFO] Start loades command as `upsert`")
	case "reindex":
		logger.Printf("[INFO] start loades command as `reindex`")
		rc, err = cmd.Reindex.Do(ctx, *name, index)
	default:
		logger.Printf("[ERROR] unknown command `%s`", strings.Join(args, "` `"))
		return 3
	}

	if err != nil {
		logger.Printf("[ERROR] %s", err.Error())
		return 4
	}
	// ok!
	logger.Printf("[INFO] %v", rc.Values()) // logger.Printf("[INFO] %#+v", rc.Values())

	return 0
}

func initInject(env util.Environment, db *sql.DB, esc *elastic.Client, cmd *commander) {
	// Injects dependecies
	var g inject.Graph
	var rt = echo.New()

	err := g.Provide(
		&inject.Object{Value: env},
		&inject.Object{Value: db},
		&inject.Object{Value: esc},
		&inject.Object{Value: &reindex{}},
	)
	if err != nil {
		logger.Panicf("failed to process inject.Provide: %s", err)
	}

	dataes.Inject(env, &g, rt)
	rdb.Inject(env, &g, rt)
	storage.Inject(env, &g, rt)
	repo.Inject(env, &g, rt)

	// inject to commander
	if err := g.Provide(&inject.Object{Value: cmd}); err != nil {
		logger.Panicf("failed to process inject.Provide: %s", err)
	}
	if err := g.Populate(); err != nil {
		logger.Panicf("failed to process inject.Populate: %s", err)
	}
}
