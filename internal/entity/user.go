package entity

import (
	"github.com/google/uuid"
	"reflect"
	"strings"
	"time"
)

type User struct {
	ID          uuid.UUID `json:"id,omitempty" sqlx:"id"`
	Username    string    `json:"username,omitempty" sqlx:"username"`
	Email       string    `json:"email" sqlx:"email"`
	Role        int32     `json:"role" sqlx:"role"`
	IsConfirmed bool      `json:"is_confirmed,omitempty" sqlx:"is_confirmed"`
	Password    string    `json:"password" sqlx:"password"`
	CreatedAt   time.Time `json:"created_at" sqlx:"created_at"`
}

func (u User) GetID() uuid.UUID {
	return u.ID
}

func (u User) GetTableName() string {
	return strings.ToLower(reflect.ValueOf(u).Type().Name() + "s")
}
