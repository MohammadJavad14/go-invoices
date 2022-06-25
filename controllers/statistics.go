package controllers

import (
	"net/http"

	"github.com/MohammadJavad14/go-invoices.git/models"
	"github.com/gin-gonic/gin"
)

// GET /total_paid_invoices
// Get the total paid invoices
func TotalPaidInvoices(ctx *gin.Context) {
	var paidInvoices []models.Invoice
	var sum float32
	models.DB.Where("is_paid = ?", true).Find(&paidInvoices).Select("sum(price)").Scan(&sum)
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	ctx.JSON(http.StatusOK, gin.H{"sum": sum})
}
