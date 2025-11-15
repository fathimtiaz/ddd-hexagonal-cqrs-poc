package usecase

import (
	"ddd-hexagonal-cqrs-poc/services/serviceA/domain"
	"ddd-hexagonal-cqrs-poc/services/serviceA/outbound/repo"
	serviceb "ddd-hexagonal-cqrs-poc/services/serviceA/outbound/serviceB"
)

func NewUseCases(
	writer repo.Writer,
	reader repo.Reader,
) domain.UseCases {
	return domain.UseCases{
		EntityAMoveUseCase: NewEntityAMoveUseCase(reader, writer, serviceb.NewServiceBCommands("intraservice")),
	}
}
