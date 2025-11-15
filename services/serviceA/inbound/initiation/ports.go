package initiation

import "ddd-hexagonal-cqrs-poc/services/common/io/input"

type Handler interface {
	InitEntityB(input.Context)
}
