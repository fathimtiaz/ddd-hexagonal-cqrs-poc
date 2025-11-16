package intraservice

import (
	"context"

	inputCommand "ddd-hexagonal-cqrs-poc/services/serviceA/io/input/command"
)

func (i *Invoker) EntityAMove(ctx context.Context, cmd inputCommand.EntityAMoveCommand) error {
	return i.handler.EntityAMove(ctx, cmd)
}

func (i *Invoker) EntityACreate(ctx context.Context, cmd inputCommand.EntityACreateCommand) error {
	return i.handler.EntityACreate(ctx, cmd)
}
