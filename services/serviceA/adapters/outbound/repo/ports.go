package repo

import (
	"context"
	"ddd-hexagonal-cqrs-poc/services/serviceA/domain"
	"ddd-hexagonal-cqrs-poc/services/serviceA/io/input/query"
)

type Reader interface {
	EntityA(context.Context, query.EntityAQuery) (domain.EntityA, error)
}

type Writer interface {
	EntityASave(context.Context, domain.EntityA)
}
