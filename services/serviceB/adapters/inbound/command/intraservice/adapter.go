package intraservice

import (
	"ddd-hexagonal-cqrs-poc/services/common/io/input"
	"ddd-hexagonal-cqrs-poc/services/common/registry"
	"ddd-hexagonal-cqrs-poc/services/serviceB/domain"
	"ddd-hexagonal-cqrs-poc/services/serviceB/io/input/command"
)

type adapter struct {
	useCases domain.UseCases
}

func NewAdapter() *adapter {
	return &adapter{registry.GetServiceBUseCases()}
}

func (a *adapter) EntityAHasMoved(c input.Context) {
	command, ok := c.Commander().(command.EntityAHasMovedCommand)
	if !ok {
		return
	}

	a.useCases.EntityAHasMoved.Execute(c.Context(), command)
}
