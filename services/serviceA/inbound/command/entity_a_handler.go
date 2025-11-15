package command

import (
	"context"
	"ddd-hexagonal-cqrs-poc/services/serviceA/io/input/command"
)

func (a *handler) EntityACreate(ctx context.Context, cmd command.EntityACreateCommand) error {
	return a.EntityAMoveUseCase.Execute(ctx, cmd)
}

func (a *handler) EntityAMove(ctx context.Context, cmd command.EntityAMoveCommand) error {
	return a.EntityAMoveUseCase.Execute(ctx, cmd)
}
