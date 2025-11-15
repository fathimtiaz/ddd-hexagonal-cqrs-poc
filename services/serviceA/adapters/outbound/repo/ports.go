package repo

import (
	"context"
	"poc/services/serviceA/domain"
	"poc/services/serviceA/io/input/query"
)

type Reader interface {
	EntityA(context.Context, query.EntityAQuery) (domain.EntityA, error)
}

type Writer interface {
	EntityASave(context.Context, domain.EntityA)
}
