package lifecycle

import (
	"context"
	"database/sql"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/pkg/errors"

	"github.com/ikeikeikeike/apicube/base/domain/usecase"
	pb "github.com/ikeikeikeike/apicube/rpc/pb/apicube/lifecycle"
)

type pingService struct {
	UsecaseLifecycles usecase.Lifecycles `inject:""`
}

func (p *pingService) GetPing(ctx context.Context, req *pb.Ping) (*pb.Ping, error) {
	return p.ping(req)
}

// ping makes sure any middleware connection established
func (p *pingService) ping(req *pb.Ping) (*pb.Ping, error) {
	msg, err := p.UsecaseLifecycles.Ping(req)

	if err != nil && errors.Cause(err) == sql.ErrNoRows {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return msg, nil
}
