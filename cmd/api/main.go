package main

import (
	"github.com/MohammadJavad14/go-invoices.git/controllers"
	"github.com/MohammadJavad14/go-invoices.git/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/users", controllers.GetUsers)

	models.ConnectDatabase()

	r.Run()
}
