package interfaces

import "github.com/ThreeDotsLabs/watermill/message"

type SubscriberHandlerRegistry interface {
	RegisterHandlerToRouter(r *message.Router) error
}

type HandlerMiddlewaresRegistry interface {
	GetHandlerMiddlewares() []message.HandlerMiddleware
}
