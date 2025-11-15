package registry

import (
	serviceACommand "ddd-hexagonal-cqrs-poc/services/serviceA/adapters/inbound/command"
	serviceBCommand "ddd-hexagonal-cqrs-poc/services/serviceB/adapters/inbound/command"
)

var (
	commandHandlerA serviceACommand.Handler
	commandHandlerB serviceBCommand.Handler
)

func GetCommandHandlerA() serviceACommand.Handler {
	return commandHandlerA
}

func GetCommandHandlerB() serviceBCommand.Handler {
	return commandHandlerB
}
