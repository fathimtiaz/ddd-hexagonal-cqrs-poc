package command

import (
	"context"
	"ddd-hexagonal-cqrs-poc/services/serviceA/io/input/command"
)

type Handler interface {
	EntityACreate(context.Context, command.EntityACreateCommand)
	EntityAMove(context.Context, command.EntityAMoveCommand)
}
