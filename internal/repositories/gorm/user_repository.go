package repositories

import (
	"github.com/Jackson-SM/Europa/internal/entities"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) Create(user *entities.User) *entities.User {
	if err := ur.db.Create(&user); err != nil {
		return nil
	}
	return user
}

func (ur *UserRepository) FindById(id string) *entities.User {
	var user entities.User

	if err := ur.db.Where("id", id).First(&user); err != nil {
		return nil
	}

	return &user
}
