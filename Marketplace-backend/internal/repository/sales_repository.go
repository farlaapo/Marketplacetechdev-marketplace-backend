package repository

import (
	"Marketplace-backend/internal/entity"

	"github.com/gofrs/uuid"
)

// ProductRepository represents the token repository contract
type SalesRepository interface{

	Create(sales *entity.Sales) error
	Update(sales *entity.Sales) error
	Delete(salaesID uuid.UUID) error
	Get(salesID uuid.UUID) (*entity.Sales, error)
	ListAll()  ([]*entity.Sales, error)

}