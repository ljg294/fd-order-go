package entity

// Order struct represents a row in the "order" table
type Order struct {
	SalesOrderID int64 `json:"order_id"`
	UserID       int64 `json:"user_id"`
}
