package loades

import (
	"context"
	"time"

	"github.com/facebookgo/inject"
	"github.com/ikeikeikeike/apicube/base/util"
	"github.com/robfig/cron"
)

var (
	jobs    *cron.Cron
	expires = 3 * time.Hour
)

// Inject injects dependencies
func Inject(ctx context.Context, env util.Environment, g *inject.Graph, rt interface{}) {}
