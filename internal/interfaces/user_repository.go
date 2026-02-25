package interfaces

import "github.com/Jackson-SM/Europa/internal/entities"

type UserRepositoryInterface interface {
	Create(user *entities.User) (*entities.User, string)
}
