package entity

import (
	"github.com/google/uuid"
	"reflect"
	"strings"
	"time"
)

type Comment struct {
	ID        uuid.UUID `sqlx:"id"`
	AuthorID  uuid.UUID `sqlx:"author_id"`
	ArticleID uuid.UUID `sqlx:"article_id"`
	Text      string    `sqlx:"text"`
	CreatedAt time.Time `sqlx:"created_at"`
}

func (c Comment) GetID() uuid.UUID {
	return c.ID
}

func (c Comment) GetTableName() string {
	return strings.ToLower(reflect.ValueOf(c).Type().Name())
}
