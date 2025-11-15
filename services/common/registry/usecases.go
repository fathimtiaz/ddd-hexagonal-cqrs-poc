package registry

import (
	serviceBDomain "poc/services/serviceB/domain"
)

var (
	serviceBUsecases serviceBDomain.UseCases
)

func GetServiceBUseCases() serviceBDomain.UseCases {
	return serviceBUsecases
}
