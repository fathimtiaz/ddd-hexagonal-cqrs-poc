package intraservice

import (
	"context"
	"poc/services/common/io/input"
	"poc/services/common/registry"
	serviceBCommand "poc/services/serviceB/adapters/inbound/command"
	serviceBCmd "poc/services/serviceB/io/input/command"
)

type adapter struct {
	handler serviceBCommand.Handler
}

func NewAdapter() *adapter {
	return &adapter{registry.GetCommandHandlerB()}
}

func (a *adapter) EntityAHasMoved(ctx context.Context, cmd serviceBCmd.EntityAHasMovedCommand) error {
	a.handler.EntityAHasMoved(input.WithCommand(ctx, cmd))
	return nil
}
