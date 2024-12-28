package service

import (
	"Marketplace-backend/internal/entity"
	"Marketplace-backend/internal/repository"
	"log"
	"time"

	"github.com/gofrs/uuid"
)

// Service is a interface
type ProductService interface {
	// create a new product
	CreateProduct(name, description string, price float64, stock int, Category, SKU string, ImageURLs []string, Discount float64, IsActive bool, Tags []string, AditionalInfo string) (*entity.Product, error)
	// get all products
	GetAllProducts() ([]*entity.Product, error)
	// get product by id
	GetProductById(productID uuid.UUID) (*entity.Product, error)
	// update product
	UpdateProduct(product entity.Product) error
	// delete product
	DeleteProduct(productID uuid.UUID) error
}

// productService is a struct
type productService struct {
	repo      repository.ProductRepository
	tokenRepo repository.TokenRepository
}

// CreateProduct implements ProductService.
func (s *productService) CreateProduct(name string, description string, price float64, stock int, Category string, SKU string, ImageURLs []string, Discount float64, IsActive bool, Tags []string, AditionalInfo string) (*entity.Product, error) {
	// Generate UUID
	NewUUID, err := uuid.NewV4()
	if err != nil {
		log.Printf("Error generating UUID: %v", err)
		return nil, err
	}

	product := &entity.Product{
		ID:             NewUUID,
		Name:           name,
		Description:    description,
		Price:          price,
		Stock:          stock,
		Category:       Category,
		SKU:            SKU,
		ImageURLs:      ImageURLs,
		Discount:       Discount,
		IsActive:       IsActive,
		Tags:           Tags,
		AdditionalInfo: AditionalInfo,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	if err := s.repo.Create(product); err != nil {
		return nil, err
	}

	return product, nil

}

// DeleteProduct implements ProductService.
func (s *productService) DeleteProduct(productID uuid.UUID) error {
	_, err := s.repo.Get(productID)
	if err != nil {
		return err
	}

	if err := s.repo.Delet(productID); err != nil {
		return err
	}
	return nil

}

// GetAllProducts implements ProductService.
func (s *productService) GetAllProducts() ([]*entity.Product, error) {
	product, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	return product, nil

}

// GetProductById implements ProductService.
func (s *productService) GetProductById(productID uuid.UUID) (*entity.Product, error) {
	product, err := s.repo.Get(productID)
	if err != nil {
		return nil, err
	}

	return product, nil
}

// UpdateProduct implements ProductService.
func (s *productService) UpdateProduct(product entity.Product) error {
	_, err := s.repo.Get(product.ID)
	if err != nil {
		return err
	}

	if err := s.repo.Update(&product); err != nil {
		return err
	}
	return nil
}

func NewProductService(productRepo repository.ProductRepository, tokenRepo repository.TokenRepository) ProductService {
	return &productService{
		repo:      productRepo,
		tokenRepo: tokenRepo,
	}
}
