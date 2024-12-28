package entity

import (
	"time"

	"github.com/gofrs/uuid"
)

type ReviewRating struct {
	ID        uuid.UUID `json:"id"`                            // Unique review ID
	UserID    uuid.UUID `json:"user_id" binding:"required"`    // ID of the user who wrote the review
	ProductID uuid.UUID `json:"product_id" binding:"required"` // ID of the product being reviewed
	Rating    int       `json:"rating" binding:"required"`     // Rating (1-5 scale)
	Comment   string    `json:"comment"`                       // User's comment
	CreatedAt time.Time `json:"created_at"`                    // Creation timestamp
	UpdatedAt time.Time `json:"updated_at"`                    // Last update timestamp
}
