package command

import "ddd-hexagonal-cqrs-poc/services/common/io/input"

type Handler interface {
	EntityACreate(input.Context)
	EntityAMove(input.Context)
}
