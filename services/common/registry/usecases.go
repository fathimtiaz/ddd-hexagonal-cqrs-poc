package registry

import (
	serviceBDomain "ddd-hexagonal-cqrs-poc/services/serviceB/domain"
)

var (
	serviceBUsecases serviceBDomain.UseCases
)

func GetServiceBUseCases() serviceBDomain.UseCases {
	return serviceBUsecases
}
