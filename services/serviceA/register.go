package servicea

import (
	serviceACommand "poc/services/serviceA/adapters/inbound/command"
	"poc/services/serviceA/adapters/outbound/repo"
	serviceAUseCase "poc/services/serviceA/usecase"
)

func Register(writer repo.Writer, reader repo.Reader) {
	serviceAUseCases := serviceAUseCase.NewUseCases(writer, reader)

	_ = serviceACommand.NewCommandHandler("intraservice", serviceAUseCases)
}
