package domain

import (
	"context"
	"ddd-hexagonal-cqrs-poc/services/common/io/input"
)

type UseCases struct {
	EntityAHasMoved UseCaser
}

type UseCaser interface {
	Execute(context.Context, input.Commander)
}
