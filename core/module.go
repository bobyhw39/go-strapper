package core

import (
	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc"
)

type Module struct {
	options    ModuleOptions
	Router     chi.Router
	GrpcServer grpc.Server
}

type ModuleOptions struct {
	ServiceName    string
	ServiceVersion string
	HttpAddress    string
	GrpcAddress    *string
}

func DefaultModuleOptions() ModuleOptions {
	return ModuleOptions{
		ServiceName:    "service-noname",
		ServiceVersion: "latest",
		HttpAddress:    ":8080",
	}
}

func MakeModule(opts ModuleOptions) func() Module {
	return func() Module {
		a := Module{
			options: opts,
		}
		return a
	}
}

func (a *Module) init() {
	a.startWebServer()
	a.startGRPCServer()
}

func DefaultModuleOptionsWithSetters(setters ...ModuleOptsSetter) ModuleOptions {
	options := DefaultModuleOptions()
	options.applySetters(setters...)
	return options
}

func (o *ModuleOptions) applySetters(setters ...ModuleOptsSetter) {
	for _, setToOptions := range setters {
		setToOptions(o)
	}
}

type ModuleOptsSetter func(options *ModuleOptions)
