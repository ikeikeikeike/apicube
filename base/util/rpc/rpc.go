package rpc

import (
	"database/sql"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pkg/errors"

	"github.com/ikeikeikeike/apicube/base/data/repo"
)

type (
	// Mux uses grpc mux
	Mux struct {
		GrpcMux      *grpc.Server
		GrpcEndpoint string
		GwMux        *runtime.ServeMux
		GwOpts       []grpc.DialOption
	}
)

// GRPCError returns relevant grpc error
func GRPCError(err error) error {
	if err != nil && errors.Cause(err) == sql.ErrNoRows {
		return status.Error(codes.NotFound, err.Error())
	}
	if err != nil && errors.Cause(err) == repo.ErrExists {
		return status.Error(codes.AlreadyExists, err.Error())
	}
	if err != nil {
		return status.Error(codes.InvalidArgument, err.Error())
	}

	return err
}
