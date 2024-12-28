package entity

import (
	"time"

	"github.com/gofrs/uuid"
)

// Order tracking
type Order struct {
	ID         uuid.UUID `json:"id"`
	UserID     uuid.UUID `json:"user_id"`
	ProductID  uuid.UUID `json:"product_id"`
	Quantity   int       `json:"quantity"`
	Status     string    `json:"status" binding:"required"` // Validation enforced here
	OrderedAt  time.Time `json:"ordered_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

