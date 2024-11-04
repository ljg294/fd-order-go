package service

import (
	"github.com/ljg294/fd-order/domain/entity"
	"github.com/ljg294/fd-order/domain/repository"
)

// OrderService contains business logic for handling orders
type OrderService struct {
	OrderRepo *repository.OrderRepository
}

// NewOrderService creates a new instance of OrderService
func NewOrderService(orderRepo *repository.OrderRepository) *OrderService {
	return &OrderService{OrderRepo: orderRepo}
}

// GetOrderByID retrieves an order by ID and handles the business logic
func (s *OrderService) GetOrderByID(salesOrderID int64) (*entity.Order, error) {
	order, err := s.OrderRepo.GetOrder(salesOrderID)
	if err != nil {
		return nil, err
	}

	return order, nil
}
