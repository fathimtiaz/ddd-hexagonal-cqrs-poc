package intraservice

import (
	inboundCommand "ddd-hexagonal-cqrs-poc/services/serviceA/inbound/command"
)

// Invoker is the cross-service bridge for in-process calls to ServiceA from other services.
type Invoker struct {
	handler inboundCommand.Handler
}

// NewInvokerWithHandler constructs an inbound invoker with injected handler.
func NewInvokerWithHandler(handler inboundCommand.Handler) *Invoker {
	return &Invoker{handler}
}
