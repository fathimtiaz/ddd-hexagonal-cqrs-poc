package intraservice

import (
	"poc/services/common/io/input"
	"poc/services/common/registry"
	"poc/services/serviceB/domain"
	"poc/services/serviceB/io/input/command"
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
