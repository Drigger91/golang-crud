package models

import (
	"time"
)

type Transaction struct {
	ID        uint    `gorm:"primaryKey"`
	UserID    uint    `gorm:"not null"`
	Amount    float64 `gorm:"type:numeric;not null"`
	CreatedAt time.Time
}
