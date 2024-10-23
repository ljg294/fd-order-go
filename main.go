package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

// Order struct represents a row in the "order" table
type Order struct {
	SalesOrderID int64 `json:"order_id"`
	UserID       int64 `json:"user_id"`
}

// Initialize MySQL connection
func setupDB() (*sql.DB, error) {
	// Use your own credentials and database details
	dsn := "root:root@tcp(localhost:22201)/local_order?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Verify connection to the database
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to the database")
	return db, nil
}

func getOrder(db *sql.DB, salesOrderID int64) (*Order, error) {
	// Corrected query without unnecessary backslashes
	fmt.Println("Query started")
	query := `SELECT sales_order_id, user_id FROM sales_order WHERE sales_order_id = ? AND delete_yn = 'N'`

	row := db.QueryRow(query, salesOrderID)
	fmt.Println("Query executed")
	var order Order
	err := row.Scan(&order.SalesOrderID, &order.UserID)
	fmt.Println("Scanned")
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &order, nil
}

func main() {
	// Initialize Gin
	router := gin.Default()

	// Connect to MySQL
	db, err := setupDB()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer db.Close()

	// GET /v1/orders/:order_id
	router.GET("/v1/orders/:sales_order_id", func(c *gin.Context) {
		// Get the order_id from the request URL
		salesOrderID := c.Param("sales_order_id")

		// Convert the orderID to int64
		var id int64
		_, err := fmt.Sscan(salesOrderID, &id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
			return
		}

		// Fetch the order from the database
		order, err := getOrder(db, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
			return
		}

		// If no order found, return 404
		if order == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
			return
		}

		// Return the order in JSON format
		c.JSON(http.StatusOK, order)
	})

	// Start the server
	router.Run(":8080")
}
