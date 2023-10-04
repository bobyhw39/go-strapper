package core

import "github.com/go-chi/chi/v5"

type Module struct {
	options ModuleOptions
	Router  chi.Router
}

type ModuleOptions struct {
	ServiceName    string
	ServiceVersion string
	HttpAddress    string
	Configuration  ModuleOptsConfig
}

func DefaultModuleOptions() ModuleOptions {
	return ModuleOptions{
		ServiceName:    "service-noname",
		ServiceVersion: "latest",
		HttpAddress:    ":8080",
		Configuration: ModuleOptsConfig{
			Path:              "environment/config.yaml",
			EnvironmentPrefix: "newservice",
			RootKey:           "newservice",
		},
	}
}

type ModuleOptsConfig struct {
	Path              string
	EnvironmentPrefix string
	RootKey           string
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
