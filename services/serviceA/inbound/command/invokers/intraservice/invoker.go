package intraservice

import (
	"context"
	inboundCommand "ddd-hexagonal-cqrs-poc/services/serviceA/inbound/command"
	inputCommand "ddd-hexagonal-cqrs-poc/services/serviceA/io/input/command"
)

// Invoker is the cross-service bridge for in-process calls to ServiceA from other services.
type Invoker struct {
	handler inboundCommand.Handler
}

// NewInvokerWithHandler constructs an inbound invoker with injected handler.
func NewInvokerWithHandler(handler inboundCommand.Handler) *Invoker {
	return &Invoker{handler}
}

func (i *Invoker) EntityAMove(ctx context.Context, cmd inputCommand.EntityAMoveCommand) error {
	return i.handler.EntityAMove(ctx, cmd)
}

func (i *Invoker) EntityACreate(ctx context.Context, cmd inputCommand.EntityACreateCommand) error {
	return i.handler.EntityACreate(ctx, cmd)
}
