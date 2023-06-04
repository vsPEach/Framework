package entity

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID          uuid.UUID `json:"id,omitempty"`
	Username    string    `json:"username,omitempty"`
	Email       string    `json:"email"`
	Role        int32     `json:"role"`
	IsConfirmed bool      `json:"is_confirmed,omitempty"`
	Password    string    `json:"password"`
	CreatedAt   time.Time `json:"created_at"`
}

func (u User) GetID() uuid.UUID {
	return u.ID
}

func (u User) GetTableName() string {
	return "users"
}
