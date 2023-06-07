package entity

import (
	"github.com/google/uuid"
	"reflect"
	"strings"
	"time"
)

type Articles []Article

type Article struct {
	ID        uuid.UUID `sqlx:"id" db:"id"`
	AuthorID  uuid.UUID `sqlx:"author_id" db:"author_id"`
	Title     string    `sqlx:"title" db:"title"`
	Text      string    `sqlx:"text" db:"text"`
	CreatedAt time.Time `sqlx:"created_at" db:"created_at"`
}

func (a Article) GetID() uuid.UUID {
	return a.ID
}

func (a Article) GetTableName() string {
	return strings.ToLower(reflect.ValueOf(a).Type().Name() + "s")
}
