package domain

import (
	"context"
	"poc/services/common/io/input"
)

type UseCases struct {
	EntityAMoveUseCase UseCaser
}

type UseCaser interface {
	Execute(context.Context, input.Commander) error

	// EntityACreate(context.Context, command.EntityACreateCommand)
	// EntityAMove(context.Context, command.EntityAMoveCommand)
}

type Publisher interface {
	EntityAHasMoved()
}
