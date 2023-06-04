package entity

import "github.com/google/uuid"

type Models interface {
	GetTableName() string
	GetID() uuid.UUID
}
