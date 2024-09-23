package models

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"type:varchar(100);not null"`
	Email     string `gorm:"unique;type:varchar(100);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
