package gateway

import (
	"Marketplace-backend/internal/entity"
	"Marketplace-backend/internal/repository"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/gofrs/uuid"
)

// SalesManagementRepository struct
type SalesRepository struct {
	db *sql.DB
}

// Create implements repository.SalesManagementRepository.
func (r *SalesRepository) Create(sales *entity.Sales) error {
	// generete UUID
	newUUD, err := uuid.NewV4()
	if err != nil {
		log.Printf("ERROR uuid : %v", err)
		return err
	}
	sales.ID = newUUD

	// query
	query := `
    INSERT INTO sales (id, user_id, total_sales, total_orders, order_id, product_id, quantity, total_price, order_data, status, created_at)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
    RETURNING id
`

	err = r.db.QueryRow(query, sales.ID, sales.UserID, sales.TotalSales, sales.TotalOrders, sales.OrdeID, sales.ProductID, sales.Quantity, sales.TotalPrice, sales.OrderDate, sales.Status, time.Now()).Scan(&sales.ID)
	if err != nil {
		log.Printf("Error creating sales: %v", err)
		return err
	}

	log.Printf("Product created with ID: %s", sales.ID)
	return nil

}

// Delete implements repository.SalesManagementRepository.
func (r *SalesRepository) Delete(salaesID uuid.UUID) error {
	query := "DELETE FROM sales WHERE id = $1"
	result, err := r.db.Exec(query, salaesID)
	if err != nil {
		log.Printf(" Error deleting sales: %v", err)
		return err
	}

	rowsAffeceted, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error getting rows affected: %v", err)
		return nil
	}

	if rowsAffeceted == 0 {
		log.Printf(" No rows effected :")
		return nil
	}

	log.Printf(" Rows affected: %d", rowsAffeceted)
	return nil

}

// Get implements repository.SalesManagementRepository.
func (r *SalesRepository) Get(salesID uuid.UUID) (*entity.Sales, error) {
	// Define struct
	var sales entity.Sales

	query := "SELECT id, user_id, total_sales, total_orders, order_id, product_id, quantity, total_price, order_data, status, created_at, updated_at FROM sales WHERE id = $1"
	err := r.db.QueryRow(query, salesID).Scan(
		&sales.ID,
		&sales.UserID,
		&sales.TotalSales,
		&sales.TotalOrders,
		&sales.OrdeID,
		&sales.ProductID,
		&sales.Quantity,
		&sales.TotalPrice, // Now it will scan correctly as float64
		&sales.OrderDate,
		&sales.Status,
		&sales.CreatedAt,
		&sales.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("sales not found")
			return nil, fmt.Errorf("no sales found")
		}

		log.Printf("Error querying product: %v", err)
		return nil, err
	}
	return &sales, nil
}

// List implements repository.SalesManagementRepository.
func (r *SalesRepository) ListAll() ([]*entity.Sales, error) {
	// Query
	row, err := r.db.Query(
		`SELECT id, user_id, total_sales, total_orders, order_id, product_id, quantity, total_price, order_data, status, created_at, updated_at FROM sales`)
	if err != nil {
		log.Printf("Error querying sales: %v", err)
		return nil, err
	}

	defer row.Close()
	var sales []*entity.Sales
	for row.Next() {
		var sale entity.Sales
		err := row.Scan(
			&sale.ID,
			&sale.UserID,
			&sale.TotalSales,
			&sale.TotalOrders,
			&sale.OrdeID,
			&sale.ProductID,
			&sale.Quantity,
			&sale.TotalPrice,
			&sale.OrderDate, // This should match the column name 'order_data'
			&sale.Status,
			&sale.CreatedAt,
			&sale.UpdatedAt,
		)
		if err != nil {
			log.Printf("Error scanning sales row: %v", err)
			return nil, err
		}
		sales = append(sales, &sale)
	}
	if err := row.Err(); err != nil {
		log.Printf("Error iterating over sales rows: %v", err)
		return nil, err
	}
	return sales, nil
}

// Update implements repository.SalesManagementRepository.
func (r *SalesRepository) Update(sales *entity.Sales) error {
	query := `UPDATE sales
	SET 
		user_id = $1, 
		total_sales = $2, 
		total_orders = $3, 
		order_id = $4, 
		product_id = $5, 
		quantity = $6, 
		total_price = $7, 
		order_data = $8, 
		status = $9, 
		created_at = $10, 
		updated_at = $11
	WHERE id = $12`

	// Execute the query
	result, err := r.db.Exec(query,
		sales.UserID,
		sales.TotalSales,
		sales.TotalOrders,
		sales.OrdeID,
		sales.ProductID,
		sales.Quantity,
		sales.TotalPrice,
		sales.OrderDate,
		sales.Status,
		time.Now(),
		time.Now(),
		sales.ID,
	)
	if err != nil {
		log.Printf("Error updating sales: %v", err)
		return err
	}

	// Check rows affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error getting rows affected: %v", err)
		return err
	}
	if rowsAffected == 0 {
		log.Printf("No rows affected by update")
		return fmt.Errorf("no rows updated")
	}

	log.Printf("Updated sales with ID %s", sales.ID)
	return nil
}

// NewSalesManagementRepository returns a new instance SalesManagementRepository
func NewSalesRepository(db *sql.DB) repository.SalesRepository {
	return &SalesRepository{db: db}

}
