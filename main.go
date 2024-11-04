package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ljg294/fd-order/app/controller"
	"github.com/ljg294/fd-order/domain/config"
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
