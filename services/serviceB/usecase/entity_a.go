package service

import (
	"context"
	"ddd-hexagonal-cqrs-poc/services/common/io/input"
	"ddd-hexagonal-cqrs-poc/services/serviceA/outbound/repo"
)

type EntityAHasMovedUseCase struct {
	reader repo.Reader
	writer repo.Writer
}

func NewEntityAHasMovedUseCase(r repo.Reader, w repo.Writer) *EntityAHasMovedUseCase {
	return &EntityAHasMovedUseCase{reader: r, writer: w}
}

func (s *EntityAHasMovedUseCase) Execute(ctx context.Context, commander input.Commander) {

}
