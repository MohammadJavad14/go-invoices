package main

import (
	"github.com/MohammadJavad14/go-invoices.git/controllers"
	"github.com/MohammadJavad14/go-invoices.git/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/users", controllers.GetUsers)
	r.GET("/users/:id", controllers.GetUser)
	r.POST("/users", controllers.CreateUser)
	r.PATCH("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

	r.GET("/users/:id/invoices", controllers.GetInvoices)
	r.GET("/users/:id/invoices/:invoice_id", controllers.GetInvoice)
	r.POST("/users/:id/invoices", controllers.CreateInvoice)
	r.PATCH("/users/:id/invoices/:invoice_id", controllers.UpdateInvoice)
	r.DELETE("/users/:id/invoices/:invoice_id", controllers.DeleteInvoice)

	models.ConnectDatabase()

	r.Run()
}
