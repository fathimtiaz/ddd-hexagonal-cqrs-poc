package intraservice

import (
	"ddd-hexagonal-cqrs-poc/services/common/io/input"
	"ddd-hexagonal-cqrs-poc/services/serviceA/io/input/command"
)

func (a adapter) EntityACreate(c input.Context) {
	cmd, ok := c.Commander().(command.EntityACreateCommand)
	if !ok {
		return
	}

	a.entityACreateUseCase.Execute(c.Context(), cmd)
}

func (a adapter) EntityAMove(c input.Context) {
	cmd, ok := c.Commander().(command.EntityAMoveCommand)
	if !ok {
		return
	}

	a.entityAMoveUseCase.Execute(c.Context(), cmd)
}
