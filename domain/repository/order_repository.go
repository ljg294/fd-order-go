package repository

import (
	"database/sql"
	"fd-order/domain/entity"
)

// OrderRepository handles database operations related to the "Order" entity
type OrderRepository struct {
	DB *sql.DB
}

// NewOrderRepository creates a new instance of OrderRepository
func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

// GetOrder fetches an order by ID from the database
func (repo *OrderRepository) GetOrder(salesOrderID int64) (*entity.Order, error) {
	query := `SELECT sales_order_id, user_id FROM sales_order WHERE sales_order_id = ? AND delete_yn = 'N'`
	row := repo.DB.QueryRow(query, salesOrderID)

	var order entity.Order
	err := row.Scan(&order.SalesOrderID, &order.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &order, nil
}
