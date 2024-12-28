package service

import (
	"Marketplace-backend/internal/entity"
	"Marketplace-backend/internal/repository"
	"log"
	"time"

	"github.com/gofrs/uuid"
)

type SalesService interface {
	// create
	CreateSales(UserID, OrdeID, ProductID uuid.UUID, TotalSales, TotalPrice float64, TotalOrders, Quantity int, OrderDate time.Time, Status string) (*entity.Sales, error)
	// Updata sales
	UpdateSales(sales entity.Sales) error
	// Delete sales
	DeleteSales(salesID uuid.UUID) error
	// GET by ID
	GetSalesByID(salesID uuid.UUID) (*entity.Sales, error)
	// GetALL SALES
	GetAllSales() ([]*entity.Sales, error)
}

// struct
type salesService struct {
	repo      repository.SalesRepository
	tokenRepo repository.TokenRepository
}

// CreateSales implements SalesService.
func (s *salesService) CreateSales(UserID, OrdeID, ProductID uuid.UUID, TotalSales, TotalPrice float64, TotalOrders, Quantity int, OrderDate time.Time, Status string) (*entity.Sales, error) {
	// generete UUID
	newUUD, err := uuid.NewV4()
	if err != nil {
		log.Printf(" Error generating UUID: %v", err)
		return nil, err
	}

	sales := entity.Sales{
		ID:          newUUD,
		UserID:      UserID,
		OrdeID:      OrdeID,
		ProductID:   ProductID,
		TotalSales:  TotalSales,
		TotalPrice:  TotalPrice,
		TotalOrders: TotalOrders,
		Quantity:    Quantity,
		OrderDate:   time.Now(),
		Status:      Status,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	if err := s.repo.Create(&sales); err != nil {
		return nil, err
	}
	return &sales, nil

}

// DeleteSales implements SalesService.
func (s *salesService) DeleteSales(salesID uuid.UUID) error {
	_, err := s.repo.Get(salesID)
	if err != nil {
		return err
	}

	if err := s.repo.Delete(salesID); err != nil {
		return err
	}

	return nil
}

// GetAllSales implements SalesService.
func (s *salesService) GetAllSales() ([]*entity.Sales, error) {
	sales, err := s.repo.ListAll()
	if err != nil {
		return nil, err
	}
	return sales, nil
}

// GetSalesByID implements SalesService.
func (s *salesService) GetSalesByID(salesID uuid.UUID) (*entity.Sales, error) {
	sales, err := s.repo.Get(salesID)
	if err != nil {
		return nil, err
	}
	return sales, nil
}

// UpdateSales implements SalesService.
func (s *salesService) UpdateSales(sales entity.Sales) error {
	_, err := s.repo.Get(sales.ID)
	if err != nil {
		return err
	}
	if err := s.repo.Update(&sales); err != nil {
		return err
	}
	return nil
}

// New salesService retun instance
func NewSalesService(salesRepo repository.SalesRepository, tokenRepo repository.TokenRepository) SalesService {
	return &salesService{
		repo:      salesRepo,
		tokenRepo: tokenRepo,
	}
}
