package command

import "ddd-hexagonal-cqrs-poc/services/common/io/input"

type EntityAHasMovedCommand struct {
}

func (c EntityAHasMovedCommand) ValidateQuery() error {
	return nil
}

func (c EntityAHasMovedCommand) ValidatePayload() error {
	return input.ValidatePayload(c)
}
