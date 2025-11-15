package intraservice

import (
	"context"

	serviceBIntraservice "ddd-hexagonal-cqrs-poc/services/serviceB/inbound/command/invokers/intraservice"
	serviceBCmd "ddd-hexagonal-cqrs-poc/services/serviceB/io/input/command"
)

type adapter struct {
	invoker *serviceBIntraservice.Invoker
}

func NewAdapter() *adapter {
	return &adapter{
		invoker: serviceBIntraservice.NewInvoker(),
	}
}

func (a *adapter) EntityAHasMoved(ctx context.Context, cmd serviceBCmd.EntityAHasMovedCommand) error {
	return a.invoker.EntityAHasMoved(ctx, cmd)
}
