package database

import (
	"github.com/Jackson-SM/Europa/internal/domain"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
	return db.AutoMigrate(
		&domain.User{},
	)
}
