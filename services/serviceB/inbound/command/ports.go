package command

import (
	"context"
	"ddd-hexagonal-cqrs-poc/services/serviceB/io/input/command"
)

type Handler interface {
	EntityAHasMoved(context.Context, command.EntityAHasMovedCommand)
}
