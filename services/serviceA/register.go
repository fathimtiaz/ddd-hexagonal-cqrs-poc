package servicea

import (
	"ddd-hexagonal-cqrs-poc/services/common/registry"
	"ddd-hexagonal-cqrs-poc/services/serviceA/outbound/repo"

	serviceACommand "ddd-hexagonal-cqrs-poc/services/serviceA/inbound/command"
	serviceAUseCase "ddd-hexagonal-cqrs-poc/services/serviceA/usecase"
)

func Register(writer repo.Writer, reader repo.Reader) {
	serviceAUseCases := serviceAUseCase.NewUseCases(writer, reader)

	serviceACommandHandler := serviceACommand.NewHandler(serviceAUseCases)

	registry.SetCommandHandlerA(serviceACommandHandler)
}
