package core

import (
	"context"
	"github.com/bobyhw39/go-strapper/interfaces"
)

type Application interface {
	AddHTTPHandlerRegistry(registry interfaces.HTTPHandlerRegistry) error
	AddHTTPMiddlewareRegistry(registry interfaces.HTTPMiddlewareRegistry) error
	ApplyRegistries() error
	Run(ctx context.Context) error
}

type Runner interface {
	Run(context context.Context) error
}
