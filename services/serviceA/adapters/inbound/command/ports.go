package command

import "poc/services/common/io/input"

type Handler interface {
	EntityACreate(input.Context)
	EntityAMove(input.Context)
}
