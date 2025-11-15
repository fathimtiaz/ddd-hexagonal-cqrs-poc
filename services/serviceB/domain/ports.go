package domain

import (
	"context"
	"poc/services/common/io/input"
)

type UseCases struct {
	EntityAHasMoved UseCaser
}

type UseCaser interface {
	Execute(context.Context, input.Commander)
}
