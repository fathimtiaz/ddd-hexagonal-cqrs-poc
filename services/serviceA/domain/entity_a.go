package domain

import "github.com/google/uuid"

type EntityA struct {
	ID        int64
	UUID      uuid.UUID
	EntityBID int64
	Name      string
	Position  string
	Frozen    bool

	EntityB *EntityB
}

func (a EntityA) CanMove() bool {
	return !a.Frozen
}

func (a *EntityA) Move(direction string) {
	a.Position = direction
}
