package handler

import (
	"google.golang.org/grpc"

	"github.com/facebookgo/inject"

	"github.com/ikeikeikeike/apicube/base/util"
	"github.com/ikeikeikeike/apicube/rpc/handler/lifecycle"
)

// Inject injects dependencies
func Inject(ctx context.Context, env util.Environment, g *inject.Graph, rt *grpc.Server) {
	product.Inject(ctx, env, g, rt)
	lifecycle.Inject(ctx, env, g, rt)
}
