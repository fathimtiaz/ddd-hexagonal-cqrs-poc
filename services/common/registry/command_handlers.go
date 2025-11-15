package registry

import (
	serviceACommand "ddd-hexagonal-cqrs-poc/services/serviceA/inbound/command"
	serviceBCommand "ddd-hexagonal-cqrs-poc/services/serviceB/inbound/command"
)

var (
	commandHandlerA serviceACommand.Handler
	commandHandlerB serviceBCommand.Handler
)

func SetCommandHandlerA(handler serviceACommand.Handler) {
	commandHandlerA = handler
}

func GetCommandHandlerA() serviceACommand.Handler {
	return commandHandlerA
}

func SetCommandHandlerB(handler serviceBCommand.Handler) {
	commandHandlerB = handler
}

func GetCommandHandlerB() serviceBCommand.Handler {
	return commandHandlerB
}
