package controllers

import (
	"net/http"

	"github.com/MohammadJavad14/go-invoices.git/models"
	"github.com/gin-gonic/gin"
)

// GET /users/:id/invoices
// Get all invoices of a user
func GetInvoices(ctx *gin.Context) {
	var user models.User
	err := models.DB.Where("id = ?", ctx.Param("id")).First(&user).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	var invoices []models.Invoice
	err = models.DB.Where("user_id = ?", ctx.Param("id")).Find(&invoices).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": invoices})

}

// GET /users/:id/invoices/:id
// Get an invoice of a user
func GetInvoice(ctx *gin.Context) {
	var user models.User
	err := models.DB.Where("id = ?", ctx.Param("id")).First(&user).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	var invoice models.Invoice
	err = models.DB.Where("user_id = ?", ctx.Param("id")).
		Where("id = ?", ctx.Param("invoice_id")).
		First(&invoice).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": invoice})

}

// POST /users/:id/invoices
// Create invoices for a user
func CreateInvoice(ctx *gin.Context) {
	var user models.User
	err := models.DB.Where("id = ?", ctx.Param("id")).First(&user).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	var input models.CreateInvoiceInput
	err = ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	invoice := models.Invoice{}
	invoice.UserID = user.ID
	invoice.Price = input.Price

	invoice.PaidAt = input.PaidAt
	models.DB.Create(&invoice)

	ctx.JSON(http.StatusOK, gin.H{"data": invoice})
}

// PATCH /users/:id/invoices/:id
// Update user invoices
func UpdateInvoice(ctx *gin.Context) {
	var user models.User
	err := models.DB.Where("id = ?", ctx.Param("id")).First(&user).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	var invoice models.Invoice
	err = models.DB.Where("user_id = ?", ctx.Param("id")).
		Where("id = ?", ctx.Param("invoice_id")).
		First(&invoice).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var input models.UpdateInvoiceInput
	err = ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&invoice).Updates(map[string]interface{}{
		"price":   input.Price,
		"is_paid": input.IsPaid,
		"paid_at": input.PaidAt,
	})

	ctx.JSON(http.StatusOK, gin.H{"data": invoice})
}

// DELETE /users/:id/invoices/:id
// Delete an user's invoice
func DeleteInvoice(ctx *gin.Context) {
	var user models.User
	err := models.DB.Where("id = ?", ctx.Param("id")).First(&user).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	var invoice models.Invoice
	err = models.DB.Where("user_id = ?", ctx.Param("id")).
		Where("id = ?", ctx.Param("invoice_id")).
		First(&invoice).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Delete(&invoice)

	ctx.JSON(http.StatusOK, gin.H{"data": true})

}
