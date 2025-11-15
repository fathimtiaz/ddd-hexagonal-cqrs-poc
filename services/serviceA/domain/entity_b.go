package domain

import "github.com/google/uuid"

type EntityB struct {
	ID   int64
	UUID uuid.UUID
	Name string
}
