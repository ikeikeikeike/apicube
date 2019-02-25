package product

import (
	"context"

	"github.com/ikeikeikeike/apicube/base/domain/usecase"
	"github.com/ikeikeikeike/apicube/base/util/logger"
	"github.com/ikeikeikeike/apicube/base/util/rpc"
	pb "github.com/ikeikeikeike/apicube/rpc/pb/apicube/product"
)

type productService struct {
	UsecaseProducts usecase.Products `inject:""`
}

func (p *productService) ListSimilars(ctx context.Context, req *pb.ListSimilarsRequest) (*pb.ListSimilarsResponse, error) {
	r, err := p.similars(ctx, req)
	if err != nil {
		logger.Errorf("%v", err)
	}

	return r, err
}

// similars makes sure any middleware connection established
func (p *productService) similars(ctx context.Context, req *pb.ListSimilarsRequest) (*pb.ListSimilarsResponse, error) {
	msg, err := p.UsecaseProducts.Similars(ctx, req)
	if err := rpc.GRPCError(err); err != nil {
		return nil, err
	}

	return msg, nil
}
