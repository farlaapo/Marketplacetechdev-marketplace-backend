package gateway

import (
	"Marketplace-backend/internal/entity"
	"Marketplace-backend/internal/repository"
	"database/sql"
	
	"log"

	"github.com/gofrs/uuid"
)

// OrderRepository represents the order repository struct
type OrderRepository struct {
	db *sql.DB
}

// Create implements repository.OrderRepository.
func (r *OrderRepository) Create(order *entity.Order) error {
	// genere UUID
	newUUID, err := uuid.NewV4()
	if err != nil {
		return err
	}

	order.ID = newUUID

	query := `INSERT INTO orders (id, user_id, product_id, quantity, status, ordered_at)
	VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;`
	result, err := r.db.Exec(query, order.ID, order.UserID, order.ProductID, order.Quantity, order.Status, order.OrderedAt)
	if err != nil {
		log.Printf("Error creating order: %v", err)
		return err
	}
	rowsAffecetd, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error getting rows affected: %v", err)
		return nil
	}
	if rowsAffecetd == 0 {
		log.Printf("No rows affected")
		return nil
	}
	return nil
}

// Delete implements repository.OrderRepository.
func (r *OrderRepository) Delete(orderID uuid.UUID) error {
	query := "DELETE FROM orders WHERE id = $1"
	result, err := r.db.Exec(query, orderID)
	if err != nil {
		log.Printf("Error deleting order: %v", err)
		return err
	}
	rowsAffeceted, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error getting rows affected: %v", err)
		return nil
	}
	if rowsAffeceted == 0 {
		log.Printf("No rows affected")
		return nil
	}
	return nil
}

// GetAll implements repository.OrderRepository.
func (r *OrderRepository) GetAll() ([]*entity.Order, error) {
	query := "SELECT id, user_id, product_id, quantity, status, ordered_at, updated_at FROM orders"
	rows, err := r.db.Query(query)
	if err != nil {
		log.Printf("Error getting orders: %v", err)
		return nil, err
	}

	defer rows.Close()
	var orders []*entity.Order
	for rows.Next() {
		order := &entity.Order{}
		err := rows.Scan(&order.ID, &order.UserID, &order.ProductID, &order.Quantity, &order.Status, &order.OrderedAt, &order.UpdatedAt)
		if err != nil {
			log.Printf("Error scanning order: %v", err)
			return nil, err
		}
		orders = append(orders, order)
		if err := rows.Err(); err != nil {
			log.Printf("Error iterating rows: %v", err)
			return nil, err
		}
	}
	return orders, nil

}

// GetByID implements repository.OrderRepository.
func (r *OrderRepository) GetByID(orderID uuid.UUID) (*entity.Order, error) {
	var order entity.Order
	query := "SELECT id, user_id, product_id, quantity, status, ordered_at, updated_at FROM orders WHERE id = $1"
	err := r.db.QueryRow(query, orderID).Scan(&order.ID, &order.UserID, &order.ProductID, &order.Quantity, &order.Status, &order.OrderedAt, &order.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No rows found")
			return nil, nil
		}
		log.Printf("Error getting order: %v", err)
		return nil, err
	}
	return &order, nil
}

// Update implements repository.OrderRepository.
func (r *OrderRepository) Update(order *entity.Order) error {
	query := "UPDATE orders SET user_id = $1, product_id = $2, quantity = $3, status = $4, ordered_at = $5, updated_at = $6 WHERE id = $7"

	// validStatuses := map[string]bool{
	// 	"Pending":   true,
	// 	"Shipped":   true,
	// 	"Delivered": true,
	// }

	// if !validStatuses[order.Status] {
	// 	return errors.New("invalid status value")
	// }

	// Execute the query with the correct parameters
	result, err := r.db.Exec(query, order.UserID, order.ProductID, order.Quantity, order.Status, order.OrderedAt, order.UpdatedAt, order.ID)
	if err != nil {
		log.Printf("Error updating order: %v", err)
		return err
	}

	// Check the number of rows affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error getting rows affected: %v", err)
		return err
	}
	if rowsAffected == 0 {
		log.Printf("No rows affected")
		return nil
	}

	return nil
}

// NewOrderRepository creates a new instance of OrderRepository
func NewOrderRepository(db *sql.DB) repository.OrderRepository {
	return &OrderRepository{
		db: db,
	}
}
