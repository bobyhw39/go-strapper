package interfaces

import (
	"context"
	"google.golang.org/grpc"
)

type GRPCHandlerRegistry interface {
	RegisterHandlerTo(srv *grpc.Server) error
}

type GRPCInterceptorRegistry interface {
	Make(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error)
}
