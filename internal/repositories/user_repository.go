package repositories

import (
	"github.com/Jackson-SM/Europa/internal/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) Create(user *domain.User) *domain.User {
	if err := ur.db.Create(&user).Error; err != nil {
		return nil
	}
	return user
}

func (ur *UserRepository) FindById(id string) *domain.User {
	var user domain.User

	if err := ur.db.Where("id", id).First(&user).Error; err != nil {
		return nil
	}

	return &user
}
