package serviceb

import (
	"ddd-hexagonal-cqrs-poc/services/common/registry"
	"ddd-hexagonal-cqrs-poc/services/serviceB/outbound/repo"

	serviceBCommand "ddd-hexagonal-cqrs-poc/services/serviceB/inbound/command"
	serviceBUseCase "ddd-hexagonal-cqrs-poc/services/serviceB/usecase"
)

func Register(writer repo.Writer, reader repo.Reader) {
	serviceBUseCases := serviceBUseCase.NewUseCases(writer, reader)

	serviceBCommandHandler := serviceBCommand.NewHandler(serviceBUseCases)

	registry.SetCommandHandlerB(serviceBCommandHandler)
}
