package command

import (
	"context"
	"ddd-hexagonal-cqrs-poc/services/serviceA/io/input/command"
)

type Handler interface {
	EntityACreate(context.Context, command.EntityACreateCommand) error
	EntityAMove(context.Context, command.EntityAMoveCommand) error
}
