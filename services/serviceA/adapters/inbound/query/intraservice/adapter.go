package intraservice

import (
	"poc/services/common/io/input"
	"poc/services/serviceA/adapters/outbound/repo"
	"poc/services/serviceA/io/input/query"
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
