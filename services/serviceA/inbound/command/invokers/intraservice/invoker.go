package intraservice

import (
	"ddd-hexagonal-cqrs-poc/services/serviceA/inbound/command"
)

type invoker struct {
	handler command.Handler
}

// NewInvoker constructs an inbound invoker with injected handler.
func NewInvoker(handler command.Handler) *invoker {
	return &invoker{handler}
}
