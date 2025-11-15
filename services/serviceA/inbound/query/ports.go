package query

import (
	"ddd-hexagonal-cqrs-poc/services/common/io/input"
)

type Handler interface {
	EntityA(input.Context)
}
