package service

import (
	"Marketplace-backend/internal/entity"
	"Marketplace-backend/internal/repository"

	"time"

	"github.com/gofrs/uuid"
)

// OrderService is a set of methods used to manipulate order
type OrderService interface {
	// CreateOrder creates a new order
	CreateOrder(UserID, ProductID uuid.UUID, Quantity int, status string) (*entity.Order, error)
	// UpdateOrder updates an existing order
	UpdateOrder(order *entity.Order) error
	// GetOrderByID gets an order by its ID
	GetOrderByID(orderID uuid.UUID) (*entity.Order, error)
	// GetOrders gets all orders
	GetAllOrders() ([]*entity.Order, error)
	// DeleteOrder deletes an order
	DeleteOrder(orderID uuid.UUID) error
}

// orderService implements OrderService struct
type orderService struct {
	repo      repository.OrderRepository
	tokenRepo repository.TokenRepository
}

// CreateOrder implements OrderService.
func (s *orderService) CreateOrder(UserID uuid.UUID, ProductID uuid.UUID, Quantity int, status string) (*entity.Order, error) {
	// generate UUID
	newUUID, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	order := &entity.Order{
		ID:        newUUID,
		UserID:    UserID,
		ProductID: ProductID,
		Quantity:  Quantity,
		Status:    status,
		OrderedAt: time.Now(),
		UpdatedAt: time.Now(),

	}

	if err := s.repo.Create(order); err != nil {
		return nil, err
	}
	return order, nil
}

// DeleteOrder implements OrderService.
func (s *orderService) DeleteOrder(orderID uuid.UUID) error {
	_, err := s.repo.GetByID(orderID)
	if err != nil {
		return err
	}
	if err := s.repo.Delete(orderID); err != nil {
		return err
	}
	return nil
}

// GetOrderByID implements OrderService.
func (s *orderService) GetOrderByID(orderID uuid.UUID) (*entity.Order, error) {
	order, err := s.repo.GetByID(orderID)
	if err != nil {
		return nil, err
	}
	return order, nil
}

// GetOrders implements OrderService.
func (s *orderService) GetAllOrders() ([]*entity.Order, error) {
	orders, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return orders, nil
}

// UpdateOrder implements OrderService.
func (s *orderService) UpdateOrder(order *entity.Order) error {
	_, err := s.repo.GetByID(order.ID)
	if err != nil {
		return err
	}
	if err := s.repo.Update(order); err != nil {
		return err
	}
	return nil
}

// NewOrderService creates a new order service
func NewOrderService(repoOrder repository.OrderRepository, tokenRepo repository.TokenRepository) OrderService {
	return &orderService{
		repo:      repoOrder,
		tokenRepo: tokenRepo,
	}
}
