package models

import (
	"time"

	"gorm.io/gorm"
)

type Invoice struct {
	gorm.Model
	UserID uint      `json:"user_id"`
	Price  uint      `json:"price"`
	PaidAt time.Time `json:"paid_at"`
}
