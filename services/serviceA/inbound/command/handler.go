package command

import "ddd-hexagonal-cqrs-poc/services/serviceA/domain"

type handler struct {
	domain.UseCases
}

func NewHandler(usecases domain.UseCases) *handler {
	return &handler{
		usecases,
	}
}
