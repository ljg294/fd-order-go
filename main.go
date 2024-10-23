package main

import (
	"fd-order/app/controller"
	"fd-order/domain/config"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// Initialize Gin
	router := gin.Default()

	// Connect to MySQL
	db, err := config.SetupDB()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer db.Close()

	// Initialize controllers
	orderController := controller.NewOrderController(db)

	// Set up routes
	router.GET("/v1/orders/:sales_order_id", orderController.GetOrder)

	// Start the server
	router.Run(":8080")
}
