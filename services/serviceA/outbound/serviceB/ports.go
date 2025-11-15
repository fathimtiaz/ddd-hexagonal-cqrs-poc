package serviceb

import (
	"context"
	serviceBCmd "ddd-hexagonal-cqrs-poc/services/serviceB/io/input/command"
)

type CommandCallers interface {
	EntityAHasMoved(ctx context.Context, cmd serviceBCmd.EntityAHasMovedCommand) error
}

type QueriyCallers interface {
}
