package handler

import (
	"context"

	"github.com/facebookgo/inject"

	"github.com/ikeikeikeike/apicube/base/util"
	"github.com/ikeikeikeike/apicube/base/util/rpc"
	"github.com/ikeikeikeike/apicube/rpc/handler/lifecycle"
)

// Inject injects dependencies
func Inject(ctx context.Context, env util.Environment, g *inject.Graph, rt *rpc.Mux) {
	lifecycle.Inject(ctx, env, g, rt)
}
