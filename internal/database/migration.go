package database

import (
	"github.com/Jackson-SM/Europa/internal/entities"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
	return db.AutoMigrate(
		&entities.User{},
	)
}
