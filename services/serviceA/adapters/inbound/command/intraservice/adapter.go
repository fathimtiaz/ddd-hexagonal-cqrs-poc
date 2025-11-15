package intraservice

import (
	"poc/services/serviceA/domain"
)

type adapter struct {
	entityACreateUseCase domain.UseCaser
	entityAMoveUseCase   domain.UseCaser
}

// NewAdapterWithUseCase constructs an inbound adapter with injected use-cases.
func NewAdapterWithUseCase(usecases domain.UseCases) *adapter {
	return &adapter{entityAMoveUseCase: usecases.EntityAMoveUseCase}
}
