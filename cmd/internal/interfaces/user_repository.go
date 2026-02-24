package interfaces

import "github.com/Jackson-SM/Europa/cmd/internal/entities"

type UserRepositoryInterface interface {
	Create(user *entities.User) (*entities.User, string)
}