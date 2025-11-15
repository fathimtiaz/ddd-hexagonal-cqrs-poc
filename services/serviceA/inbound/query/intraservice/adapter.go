package intraservice

import (
	"ddd-hexagonal-cqrs-poc/services/common/io/input"
	"ddd-hexagonal-cqrs-poc/services/serviceA/io/input/query"
	"ddd-hexagonal-cqrs-poc/services/serviceA/outbound/repo"
)

type adapter struct {
	reader repo.Reader
}

func New(reader repo.Reader) *adapter {
	return &adapter{reader}
}

func (a *adapter) EntityA(c input.Context) {
	a.reader.EntityA(c.Context(), c.Querier().(query.EntityAQuery))
}
