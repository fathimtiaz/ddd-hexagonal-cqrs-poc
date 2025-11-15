package command

import "poc/services/common/io/input"

type Handler interface {
	EntityAHasMoved(input.Context)
}
