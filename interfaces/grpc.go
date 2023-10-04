package interfaces

import (
	"google.golang.org/grpc"
)

type GRPCHandlerRegistry interface {
	RegisterHandlerTo(srv *grpc.Server) error
}
