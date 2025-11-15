package command

import (
	"context"
	"ddd-hexagonal-cqrs-poc/services/serviceB/io/input/command"
)

func (h *handler) EntityAHasMoved(ctx context.Context, cmd command.EntityAHasMovedCommand) {
	h.usecases.EntityAHasMoved.Execute(ctx, cmd)
}
