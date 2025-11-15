package service

import (
	"ddd-hexagonal-cqrs-poc/services/serviceA/outbound/repo"
	"ddd-hexagonal-cqrs-poc/services/serviceB/domain"
)

func NewUseCases(
	writer repo.Writer,
	reader repo.Reader,
) domain.UseCases {
	return domain.UseCases{
		EntityAHasMoved: NewEntityAHasMovedUseCase(reader, writer),
	}
}
