package repository

import (
	"Marketplace-backend/internal/entity"

	"github.com/gofrs/uuid"
)

// ProductRepository represents the token repository contract
type ProductRepository interface {
	Create(product *entity.Product) error
	Update(product *entity.Product) error
	Delet(productID uuid.UUID) error
	Get(productID uuid.UUID) (*entity.Product, error)
	FindAll() ([]*entity.Product, error)
}
