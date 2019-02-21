package product

import (
	"context"
	"database/sql"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/pkg/errors"

	"github.com/ikeikeikeike/apicube/base/domain/usecase"
	"github.com/ikeikeikeike/apicube/base/util/logger"
	pb "github.com/ikeikeikeike/apicube/rpc/pb/apicube/product"
)

type productService struct {
	UsecaseProducts usecase.Products `inject:""`
}

func (p *productService) ListSimilars(ctx context.Context, req *pb.ListSimilarsRequest) (*pb.ListSimilarsResponse, error) {
	r, err := p.similars(req)
	if err != nil {
		logger.Errorf("%v", err)
	}

	return r, err
}

// similars makes sure any middleware connection established
func (p *productService) similars(req *pb.ListSimilarsRequest) (*pb.ListSimilarsResponse, error) {
	msg, err := p.UsecaseProducts.Similars(req)

	if err != nil && errors.Cause(err) == sql.ErrNoRows {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return msg, nil
}
