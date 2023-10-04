package core

import (
	"context"
	"fmt"
	"github.com/bobyhw39/go-strapper/fsmutils"
	"github.com/bobyhw39/go-strapper/interfaces"
	"github.com/go-chi/chi/v5"
	"github.com/looplab/fsm"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

const (
	_WaitingForRegistriesApplied = "application_waiting_for_registries_applied"
	_RegistriesApplied           = "application_registries_applied"
	_ApplicationStarted          = "application_started"
	_ApplyRegistries             = "application_apply_registry"
	_StartApplication            = "application_start_application"
)

type ApplicationDelegate struct {
	rootRouter           chi.Router
	handlerRegistries    []interfaces.HTTPHandlerRegistry
	middlewareRegistries []interfaces.HTTPMiddlewareRegistry
	stateMachine         *fsm.FSM
	runners              []Runner
}

func NewApplicationDelegate(
	rootRouter chi.Router,
) *ApplicationDelegate {

	app := &ApplicationDelegate{
		rootRouter:           rootRouter,
		handlerRegistries:    []interfaces.HTTPHandlerRegistry{},
		middlewareRegistries: []interfaces.HTTPMiddlewareRegistry{},
	}

	app.stateMachine = fsm.NewFSM(
		_WaitingForRegistriesApplied,
		fsm.Events{
			{
				Name: _ApplyRegistries,
				Src:  []string{_WaitingForRegistriesApplied},
				Dst:  _RegistriesApplied,
			},
			{
				Name: _StartApplication,
				Src:  []string{_RegistriesApplied},
				Dst:  _ApplicationStarted,
			},
		},
		fsm.Callbacks{
			fsmutils.AsEnterState(_ApplicationStarted): func(event *fsm.Event) {
				if len(event.Args) != 1 {
					event.Cancel(fmt.Errorf("expected 1 args passed"))
					return
				}
				ctx, ok := event.Args[0].(context.Context)
				if !ok {
					event.Cancel(fmt.Errorf("args[0] expected to be context.Context"))
					return
				}
				err := app.startApplication(ctx)
				if err != nil {
					event.Cancel(err)
					return
				}
			},
			fsmutils.AsEnterState(_RegistriesApplied): func(event *fsm.Event) {
				err := app.executeApplyRegistry()
				if err != nil {
					event.Cancel(err)
					return
				}
			},
		},
	)

	return app
}

func (a *ApplicationDelegate) AddHTTPHandlerRegistry(registry interfaces.HTTPHandlerRegistry) error {
	a.handlerRegistries = append(a.handlerRegistries, registry)
	return nil
}

func (a *ApplicationDelegate) AddHTTPMiddlewareRegistry(registry interfaces.HTTPMiddlewareRegistry) error {
	a.middlewareRegistries = append(a.middlewareRegistries, registry)
	return nil
}

func (a *ApplicationDelegate) AddRunner(runner Runner) error {
	a.runners = append(a.runners, runner)
	return nil
}

func (a *ApplicationDelegate) ApplyRegistries() error {
	return a.applyEvent(_ApplyRegistries)
}

func (a *ApplicationDelegate) Run(ctx context.Context) (err error) {
	return a.applyEvent(_StartApplication, ctx)
}

func (a *ApplicationDelegate) executeApplyRegistry() (err error) {
	a.rootRouter, err = a.applyMiddlewaresToRouter(a.rootRouter)
	if err != nil {
		return err
	}
	a.rootRouter, err = a.applyRoutesToRouter(a.rootRouter)
	if err != nil {
		return err
	}

	return
}

func (a *ApplicationDelegate) startApplication(ctx context.Context) (err error) {
	errChan := make(chan error)
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	var wg sync.WaitGroup

	wg.Add(2)
	go func(errChan chan error) {
		wg.Done()
	}(errChan)

	go func(errChan chan error) {
		wg.Done()
	}(errChan)

	go a.runRunners(ctx, &wg, errChan)

	wg.Wait()

	select {
	case <-ctx.Done():
		return nil
	case <-done:
		return nil
	case err := <-errChan:
		return err
	}

}

func (a *ApplicationDelegate) applyMiddlewaresToRouter(router chi.Router) (chi.Router, error) {
	for _, registry := range a.middlewareRegistries {
		for _, m := range registry.GetMiddlewares() {
			router.Use(m)
		}
	}
	return router, nil
}

func (a *ApplicationDelegate) applyRoutesToRouter(router chi.Router) (chi.Router, error) {
	if len(a.handlerRegistries) == 0 {
		return router, fmt.Errorf("routes registries are empty")
	}
	for _, registry := range a.handlerRegistries {
		if err := registry.RegisterRoutesTo(router); err != nil {
			return router, err
		}
	}
	return router, nil
}

func (a *ApplicationDelegate) applyEvent(event string, args ...interface{}) error {
	err := a.stateMachine.Event(event, args...)
	if err != nil {
		return err
	}
	return nil
}

func (a *ApplicationDelegate) runRunners(ctx context.Context, wg *sync.WaitGroup, errChan chan error) {
	for _, runner := range a.runners {
		wg.Add(1)
		go func(runner Runner, errChan chan error) {
			if err := runner.Run(ctx); err != nil {
				log.Println(err)
				wg.Done()
				errChan <- err
			}
			wg.Done()
		}(runner, errChan)
	}
}
