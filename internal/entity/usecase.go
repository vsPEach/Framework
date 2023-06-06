package entity

import "github.com/google/uuid"

type Model interface {
	GetTableName() string
	GetID() uuid.UUID
}
