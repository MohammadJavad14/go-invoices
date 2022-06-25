package models

import (
	"time"

	"gorm.io/gorm"
)

type InvoiceInfo struct {
	Price  uint      `json:"price"`
	IsPaid *bool     `json:"is_paid"  gorm:"default:false"`
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
	IsPaid *bool     `json:"is_paid" gorm:"default:false"`
	PaidAt time.Time `json:"paid_at"`
}

type UpdateInvoiceInput struct {
	InvoiceInfo
}
