package entity

import (
	"time"

	"github.com/gofrs/uuid"
)

type Sales struct {
	ID          uuid.UUID `json:"id"`
	UserID      uuid.UUID `json:"user_id" binding:"required"`
	TotalSales  float64   `json:"total_sales" binding:"required"`
	TotalOrders int       `json:"total_orders" binding:"required"`
	OrdeID      uuid.UUID `json:"order_id" binding:"required"`
	ProductID   uuid.UUID `json:"product_id" binding:"required"`
	Quantity    int       `json:"quantity" binding:"required"`
	TotalPrice  float64   `json:"total_price" binding:"required"`
	OrderDate   time.Time `json:"orde_data"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
