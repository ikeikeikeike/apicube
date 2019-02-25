package job

import (
	"context"

	"github.com/facebookgo/inject"

	"github.com/ikeikeikeike/apicube/base/job/loades"
	"github.com/ikeikeikeike/apicube/base/util"
)

// Inject injects dependencies
func Inject(ctx context.Context, env util.Environment, g *inject.Graph, rt interface{}) {
	loades.Inject(ctx, env, g, rt)
}
