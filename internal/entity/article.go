package entity

import (
	"github.com/google/uuid"
	"reflect"
	"strings"
	"time"
)

type Article struct {
	ID        uuid.UUID
	AuthorID  uuid.UUID
	Title     string
	Text      string
	CreatedAt time.Time
}

func (a Article) GetID() uuid.UUID {
	return a.ID
}

func (a Article) GetTableName() string {
	return strings.ToLower(reflect.ValueOf(a).Type().Name())
}
