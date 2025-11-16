package command

import "ddd-hexagonal-cqrs-poc/services/serviceA/domain"

// act as a reusable service-scoped handler that can be invoked in each invoker
// omitting the need to register all invoker in the service's register
type handler struct {
	domain.UseCases
}

func NewHandler(usecases domain.UseCases) *handler {
	return &handler{
		usecases,
	}
}
