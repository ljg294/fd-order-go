package controller

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/ljg294/fd-order/app/service"
	"github.com/ljg294/fd-order/domain/repository"
	"net/http"
	"strconv"
)

// OrderController handles HTTP requests related to orders
type OrderController struct {
	OrderService *service.OrderService
}

// NewOrderController creates a new instance of OrderController
func NewOrderController(db *sql.DB) *OrderController {
	orderRepo := service.NewOrderService(
		repository.NewOrderRepository(db),
	)
	return &OrderController{OrderService: orderRepo}
}

// GetOrder handles the GET /v1/orders/:sales_order_id route
func (ctrl *OrderController) GetOrder(c *gin.Context) {
	// Get the order_id from the request URL
	salesOrderID, err := strconv.ParseInt(c.Param("sales_order_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	// Fetch the order from the service
	order, err := ctrl.OrderService.GetOrderByID(salesOrderID)
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
}
