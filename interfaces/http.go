package interfaces

import "github.com/go-chi/chi/v5"

type HTTPHandlerRegistry interface {
	RegisterRoutesTo(r chi.Router) error
}
type HTTPMiddlewareRegistry interface {
	GetMiddlewares() chi.Middlewares
}
