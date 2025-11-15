package command

import (
	"ddd-hexagonal-cqrs-poc/services/serviceB/domain"
)

type handler struct {
	usecases domain.UseCases
}

func NewHandler(usecases domain.UseCases) *handler {
	return &handler{usecases}
}
