package rpc

import (
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

type (
	Mux struct {
		GrpcMux *grpc.Server
		GwMux   *runtime.ServeMux
		GwOpts  []grpc.DialOption
	}
)
