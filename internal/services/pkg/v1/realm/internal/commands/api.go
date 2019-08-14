package commands

import "context"

// HandlerFunc describes a function implementation.
type HandlerFunc func(context.Context, interface{}) (interface{}, error)
