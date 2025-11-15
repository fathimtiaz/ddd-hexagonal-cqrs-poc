package registry

import (
	serviceACommand "poc/services/serviceA/adapters/inbound/command"
	serviceBCommand "poc/services/serviceB/adapters/inbound/command"
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
