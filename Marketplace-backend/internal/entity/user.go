package entity

import (
	"time"

	"github.com/gofrs/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username" binding:"required"`
	Password  string    `json:"password" binding:"required,min=8"` // Minimum 8 characters for password
	Email     string    `json:"email" binding:"required,email"`    // Valid email format
	FirstName string    `json:"first_name" binding:"required"`
	LastName  string    `json:"last_name" binding:"required"`
	RoleID    uuid.UUID `json:"role_id" binding:"required"` // Valid UUID
	RoleName  string    `json:"role_name" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
