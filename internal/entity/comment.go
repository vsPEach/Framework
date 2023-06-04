package entity

import (
	"github.com/google/uuid"
	"time"
)

type Comment struct {
	ID        uuid.UUID
	AuthorID  uuid.UUID
	ArticleID uuid.UUID
	Text      string
	CreatedAt time.Time
}

func (c Comment) GetID() uuid.UUID {
	return c.ID
}

func (c Comment) GetTableName() string {
	return "comments"
}
