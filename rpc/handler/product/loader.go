package product

import (
	"context"

	"github.com/facebookgo/inject"

	"github.com/ikeikeikeike/apicube/base/util"
	"github.com/ikeikeikeike/apicube/base/util/logger"
	"github.com/ikeikeikeike/apicube/base/util/rpc"
	pb "github.com/ikeikeikeike/apicube/rpc/pb/apicube/product"
)

// Inject injects dependencies
func Inject(ctx context.Context, env util.Environment, g *inject.Graph, rt *rpc.Mux) {
	var product productService

	// inject
	err := g.Provide(
		&inject.Object{Value: &product},
	)
	if err != nil {
		logger.Panicf("[PANIC] failed to process injection: %s", err)
	}

	// routes
	pb.RegisterProductServiceServer(rt.GrpcMux, &product)

	if err := pb.RegisterProductServiceHandlerFromEndpoint(ctx, rt.GwMux, rt.GrpcEndpoint, rt.GwOpts); err != nil {
		logger.Panicf("[PANIC] failed to register grpc gateway: %s", err)
	}
}
