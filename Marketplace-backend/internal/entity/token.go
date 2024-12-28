package entity

import (
	"time"

	"github.com/gofrs/uuid"
)

// token reperents a token entity
type Token struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Token     string
	ExpiredAt time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
