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

func GetUser(ctx *gin.Context) {
	var user models.User
	err := models.DB.Where("id = ?", ctx.Param("id")).First(&user).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": user})
}

// POST /users
// Create new user
func CreateUser(ctx *gin.Context) {
	var input models.CreateUserInput
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{}
	user.Name = input.Name
	user.Gender = input.Gender
	user.BirthYear = input.BirthYear
	models.DB.Create(&user)

	ctx.JSON(http.StatusOK, gin.H{"data": user})
}

// PATCH /users/:id
// Update a user
func UpdateUser(ctx *gin.Context) {
	var user models.User
	err := models.DB.Where("id = ?", ctx.Param("id")).First(&user).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not Found!"})
		return
	}

	var input models.UpdateUserInput
	err = ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&user).Updates(map[string]interface{}{
		"name":       input.Name,
		"gender":     input.Gender,
		"birth_year": input.BirthYear,
	})

	ctx.JSON(http.StatusOK, gin.H{"data": user})
}

// DELETE /users/:id
// Delete a user
func DeleteUser(ctx *gin.Context) {
	var user models.User
	err := models.DB.Where("id = ?", ctx.Param("id")).First(&user).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not Found!"})
		return
	}

	models.DB.Delete(&user)

	ctx.JSON(http.StatusOK, gin.H{"data": true})
}
