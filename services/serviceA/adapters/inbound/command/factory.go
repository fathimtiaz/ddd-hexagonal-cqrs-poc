package command

import (
	"poc/services/serviceA/adapters/inbound/command/intraservice"
	"poc/services/serviceA/domain"
)

func NewCommandHandler(adapter string, useCases domain.UseCases) Handler {
	switch adapter {
	case "intraservice":
		return intraservice.NewAdapterWithUseCase(useCases)
	// case "otherAdapter":
	// 	return &OtherAdapterCommandHandler{}
	default:
		return nil
	}
}
