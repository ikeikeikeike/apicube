package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/gigawattio/metaflector"
	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/cast"

	"github.com/ikeikeikeike/apicube/base/data/es"
	"github.com/ikeikeikeike/apicube/base/util"
	"github.com/ikeikeikeike/apicube/base/util/logger"
)

// Usage
var (
	name = flag.String("idxname", es.ProductsName, "index name. products|favorites")
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

	var env util.Environment = &Env{}
	if err := envconfig.Process("", env); err != nil {
		logger.Printf("[ERROR] failed to process env var: %s", err)
		return 2
	}
	index := es.Schema(*name)
	if index == "" {
		logger.Printf("[ERROR] A `%s` does not have index content", *name)
		return 3
	}

	return 0
}
