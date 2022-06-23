package controllers

import (
	"net/http"

	"github.com/MohammadJavad14/go-invoices.git/models"
	"github.com/gin-gonic/gin"
)

// GET /users
// Get all users
func GetUsers(ctx *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	ctx.JSON(http.StatusOK, gin.H{"data": users})

}
