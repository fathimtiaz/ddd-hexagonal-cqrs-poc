package query

import "ddd-hexagonal-cqrs-poc/services/common/io/input"

type EntityAQuery struct {
	input.BaseQuery

	WithEntityB *EntityBQuery
}

type EntityBQuery struct {
	input.BaseQuery
}
