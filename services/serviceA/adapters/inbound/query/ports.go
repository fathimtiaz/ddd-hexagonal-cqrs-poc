package query

import (
	"poc/services/common/io/input"
)

type Handler interface {
	EntityA(input.Context)
}
