package usecase

import (
	"poc/services/serviceA/adapters/outbound/repo"
	serviceb "poc/services/serviceA/adapters/outbound/serviceB"
	"poc/services/serviceA/domain"
)

func NewUseCases(
	writer repo.Writer,
	reader repo.Reader,
) domain.UseCases {
	return domain.UseCases{
		EntityAMoveUseCase: NewEntityAMoveUseCase(reader, writer, serviceb.NewServiceBCommands("intraservice")),
	}
}
