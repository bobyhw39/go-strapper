package core

import (
	"github.com/bobyhw39/go-strapper/stringutils"
	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
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
		a.init()
		return a
	}
}

func (a *Module) init() {
	router := chi.NewRouter()
	a.Router = router
	http.ListenAndServe(a.options.HttpAddress, a.Router)

	if !stringutils.IsPointerBlank(a.options.GrpcAddress) {
		lis, err := net.Listen("tcp", *a.options.GrpcAddress)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		if err := a.GrpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %s", err)
		}
	}

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
