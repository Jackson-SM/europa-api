package user

import "time"

type UserGormRepository struct {
	ID        string `gorm:"default:uuid_generate_v3()"`
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}