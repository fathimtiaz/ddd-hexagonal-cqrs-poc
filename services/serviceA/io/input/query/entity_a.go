package query

import "poc/services/common/io/input"

type EntityAQuery struct {
	input.BaseQuery

	WithEntityB *EntityBQuery
}

type EntityBQuery struct {
	input.BaseQuery
}
