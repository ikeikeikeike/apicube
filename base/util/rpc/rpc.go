package rpc

import (
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
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
