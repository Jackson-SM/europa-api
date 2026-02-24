package repositories

import (
	"github.com/Jackson-SM/Europa/cmd/internal/entities"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) Create(user *entities.User) (*entities.User, string) {
	return user, "user created successfully"
}