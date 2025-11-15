package intraservice

import (
	"context"
	"ddd-hexagonal-cqrs-poc/services/common/registry"
	inboundCommand "ddd-hexagonal-cqrs-poc/services/serviceB/inbound/command"
	"ddd-hexagonal-cqrs-poc/services/serviceB/io/input/command"
)

// Invoker is the cross-service bridge for in-process calls to ServiceB from other services.
type Invoker struct {
	handler inboundCommand.Handler
}

func NewInvoker() *Invoker {
	return &Invoker{handler: registry.GetCommandHandlerB()}
}

func (i *Invoker) EntityAHasMoved(ctx context.Context, cmd command.EntityAHasMovedCommand) error {
	return i.handler.EntityAHasMoved(ctx, cmd)
}
