package service

import (
	"ddd-hexagonal-cqrs-poc/services/serviceB/domain"
	"ddd-hexagonal-cqrs-poc/services/serviceB/outbound/repo"
)

func NewUseCases(
	writer repo.Writer,
	reader repo.Reader,
) domain.UseCases {
	return domain.UseCases{
		EntityAHasMoved: NewEntityAHasMovedUseCase(reader, writer),
	}
}
