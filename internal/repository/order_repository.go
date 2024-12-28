package repository

import (
	"Marketplace-backend/internal/entity"

	"github.com/gofrs/uuid"
)

// OrderRepository represents the order repository contract
type OrderRepository interface {
	Create(order *entity.Order) error
	Update(order *entity.Order) error
	GetAll() ([]*entity.Order, error)
	GetByID(orderID uuid.UUID) (*entity.Order, error)
	Delete(orderID uuid.UUID) error
}