package entities

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(name string, email string) *User {
	return &User{ID: uuid.NewString(), Name: name, Email: email}
}
