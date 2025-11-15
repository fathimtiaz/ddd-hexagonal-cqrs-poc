package servicea

import (
	"ddd-hexagonal-cqrs-poc/services/common/registry"
	"ddd-hexagonal-cqrs-poc/services/serviceA/outbound/repo"

	serviceACommand "ddd-hexagonal-cqrs-poc/services/serviceA/inbound/command"
	serviceAUseCase "ddd-hexagonal-cqrs-poc/services/serviceA/usecase"

	serviceBCommand "ddd-hexagonal-cqrs-poc/services/serviceB/inbound/command"
	serviceBUseCase "ddd-hexagonal-cqrs-poc/services/serviceB/usecase"
)

func Register(writer repo.Writer, reader repo.Reader) {
	serviceAUseCases := serviceAUseCase.NewUseCases(writer, reader)

	serviceACommandHandler := serviceACommand.NewHandler(serviceAUseCases)

	registry.SetCommandHandlerA(serviceACommandHandler)

	serviceBUseCases := serviceBUseCase.NewUseCases(writer, reader)

	sereviceBCommandHandler := serviceBCommand.NewHandler(serviceBUseCases)

	registry.SetCommandHandlerB(sereviceBCommandHandler)

}
