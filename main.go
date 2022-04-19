package main

import (
	"assigment-2/controllers"
	"assigment-2/database"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.NewPostgres()
	if db == nil {
		fmt.Println("Running db is fail")
	}

	ordersController := controllers.NewControllerOrders(db)

	router := gin.Default()

	router.GET("/orders", ordersController.GetOrders)

	router.POST("/orders", ordersController.CreateOrders)

	router.PUT("/orders/:id", ordersController.UpdateOrdersByID)

	router.DELETE("/orders/:id", ordersController.DeleteOrderByID)

	router.Run(":8080")
}
