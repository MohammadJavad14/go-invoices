package models

import (
	"time"

	"gorm.io/gorm"
)

type InvoiceInfo struct {
	Price  uint      `json:"price"`
	PaidAt time.Time `json:"paid_at"`
}
type Invoice struct {
	gorm.Model
	InvoiceInfo
	UserID uint `json:"user_id"`
	User   User `json:"-"`
}

type CreateInvoiceInput struct {
	Price  uint      `json:"price" binding:"required"`
	PaidAt time.Time `json:"paid_at"`
}

type UpdateInvoiceInput struct {
	InvoiceInfo
}
