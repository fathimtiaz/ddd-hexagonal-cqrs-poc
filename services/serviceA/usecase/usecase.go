package usecase

import (
	"ddd-hexagonal-cqrs-poc/services/serviceA/adapters/outbound/repo"
	serviceb "ddd-hexagonal-cqrs-poc/services/serviceA/adapters/outbound/serviceB"
	"ddd-hexagonal-cqrs-poc/services/serviceA/domain"
)

func NewUseCases(
	writer repo.Writer,
	reader repo.Reader,
) domain.UseCases {
	return domain.UseCases{
		EntityAMoveUseCase: NewEntityAMoveUseCase(reader, writer, serviceb.NewServiceBCommands("intraservice")),
	}
}
