package command

import "ddd-hexagonal-cqrs-poc/services/common/io/input"

type Handler interface {
	EntityAHasMoved(input.Context)
}
