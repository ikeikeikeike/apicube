package loades

import (
	"context"
	"time"

	"github.com/facebookgo/inject"
	"github.com/robfig/cron"

	"github.com/ikeikeikeike/apicube/base/util"
	"github.com/ikeikeikeike/apicube/base/util/graceful"
	"github.com/ikeikeikeike/apicube/base/util/logger"
)

var (
	jobs    *cron.Cron
	expires = 50 * time.Minute
)

// Inject injects dependencies
func Inject(ctx context.Context, env util.Environment, g *inject.Graph, rt interface{}) {
	var ldProd = &loadProducts{}

	// inject
	err := g.Provide(
		&inject.Object{Value: ldProd},
	)
	if err != nil {
		logger.Panicf("[PANIC] failed to process injection: %s", err)
	}

	graceful.PreHook(func() {
		jobs.Stop()
	})

	jobs = cron.New()
	jobs.Start()
	jobs.AddFunc("@every 1h", ldProd.load)

	for _, job := range jobs.Entries() {
		logger.Printf("[INFO] %#+v next:%s", job.Job, job.Next)
	}
}
