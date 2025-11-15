package servicea

import (
	serviceACommand "ddd-hexagonal-cqrs-poc/services/serviceA/adapters/inbound/command"
	"ddd-hexagonal-cqrs-poc/services/serviceA/adapters/outbound/repo"
	serviceAUseCase "ddd-hexagonal-cqrs-poc/services/serviceA/usecase"
)

func Register(writer repo.Writer, reader repo.Reader) {
	serviceAUseCases := serviceAUseCase.NewUseCases(writer, reader)

	_ = serviceACommand.NewCommandHandler("intraservice", serviceAUseCases)
}
