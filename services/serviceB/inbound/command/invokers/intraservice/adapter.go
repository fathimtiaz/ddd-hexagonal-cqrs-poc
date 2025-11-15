package intraservice

import (
	"ddd-hexagonal-cqrs-poc/services/common/registry"
	"ddd-hexagonal-cqrs-poc/services/serviceB/inbound/command"
)

type invoker struct {
	handler command.Handler
}

func NewInvoker(handler command.Handler) *invoker {
	return &invoker{handler: registry.GetCommandHandlerB()}
}
