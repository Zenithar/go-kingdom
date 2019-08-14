package commands

import "context"

// Handler describes a command handler
type Handler interface {
	Handle(ctx context.Context, req interface{}) (interface{}, error)
}

// -----------------------------------------------------------------------------

// HandlerFunc describes a function implementation.
type HandlerFunc func(context.Context, interface{}) (interface{}, error)

// Handle call the wrapped function
func (f HandlerFunc) Handle(ctx context.Context, req interface{}) (interface{}, error) {
	return f(ctx, req)
}

// -----------------------------------------------------------------------------
