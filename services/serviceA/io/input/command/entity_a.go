package command

import (
	"poc/services/common/io/input"
	"poc/services/serviceA/io/input/query"
)

type EntityACreateCommand struct {
	Query   query.EntityAQuery
	Payload EntityACreatePayload
}

func (c EntityACreateCommand) ValidateQuery() error {
	return nil
}

func (c EntityACreateCommand) ValidatePayload() error {
	return input.ValidatePayload(c)
}

type EntityACreatePayload struct {
	Name string `json:"name"`
}

type EntityAMoveCommand struct {
	Query   query.EntityAQuery
	Payload EntityAMovePayload
}

func (c EntityAMoveCommand) ValidateQuery() error {
	return nil
}

func (c EntityAMoveCommand) ValidatePayload() error {
	return input.ValidatePayload(c)
}

type EntityAMovePayload struct {
	Direction string `json:"direction"`
}
