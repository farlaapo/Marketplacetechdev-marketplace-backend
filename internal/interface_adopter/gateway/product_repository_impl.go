package gateway

import (
	"Marketplace-backend/internal/entity"
	"Marketplace-backend/internal/repository"
	"database/sql"
	"fmt"
	"log"

	"github.com/gofrs/uuid"
)

// ProductRepositoryImpl struct
type ProductRepositoryImpl struct {
	db *sql.DB
}

// Create implements repository.ProductRepository.
func (r *ProductRepositoryImpl) Create(product *entity.Product) error {
	// Generate UUID
	NewUUID, err := uuid.NewV4()
	if err != nil {
		log.Printf("Error generating UUID: %v", err)
		return err
	}
	product.ID = NewUUID

	// Query to insert product
	query := `INSERT INTO products(id, name, description, price, stock, category, sku, image_urls, discount, is_active, tags, additional_info, created_at, updated_at)
VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) RETURNING id;`

	// Insert product and return the ID
	err = r.db.QueryRow(query, product.ID, product.Name, product.Description, product.Price, product.Stock, product.Category, product.SKU, product.ImageURLsToString(), product.Discount, product.IsActive, product.TagsToString(), product.AdditionalInfo, product.CreatedAt, product.UpdatedAt).Scan(&product.ID)

	if err != nil {
		log.Printf("Error creating product: %v", err)
		return err
	}

	log.Printf("Product created with ID: %s", product.ID)
	return nil
}

// Delet implements repository.ProductRepository.
func (r *ProductRepositoryImpl) Delet(productID uuid.UUID) error {
	query := "DELETE FROM products WHERE id = $1"

	result, err := r.db.Exec(query, productID)
	if err != nil {
		log.Printf("Error updating user: %v", err)
		return err
	}
	rowsAffeceted, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error getting rows affected: %v", err)
		return nil
	}
	if rowsAffeceted == 0 {
		log.Printf(" No rows affected")
		return nil
	}
	log.Printf(" Rows affected: %d", rowsAffeceted)
	return nil

}

// FindAll implements repository.ProductRepository.
func (r *ProductRepositoryImpl) FindAll() ([]*entity.Product, error) {
	rows, err := r.db.Query(
		"SELECT id, name, description, price, stock, category, sku, image_urls, discount, is_active, tags, additional_info, created_at, updated_at FROM products")
	if err != nil {
		log.Printf("Error querying products: %v", err)
		return nil, err
	}
	defer rows.Close()

	var products []*entity.Product
	for rows.Next() {
		var product entity.Product
		var imageURLsString, tagsString string

		// Retrieve the data
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Stock,
			&product.Category,
			&product.SKU,
			&imageURLsString, // Temporary string for image URLs
			&product.Discount,
			&product.IsActive,
			&tagsString, // Temporary string for tags
			&product.AdditionalInfo,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			log.Printf("Error scanning product row: %v", err)
			return nil, err
		}

		// Convert string fields to slices
		product.StringToImageURLs(imageURLsString)
		product.StringToTags(tagsString)

		products = append(products, &product)
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		log.Printf("Error iterating over product rows: %v", err)
		return nil, err
	}

	return products, nil
}

// Get implements repository.ProductRepository.
func (r *ProductRepositoryImpl) Get(productID uuid.UUID) (*entity.Product, error) {
	// Define product
	var product entity.Product
	var imageURLsString, tagsString string

	// Query
	query := "SELECT id, name, description, price, stock, category, sku, image_urls, discount, is_active, tags, additional_info, created_at, updated_at FROM products WHERE id = $1"
	err := r.db.QueryRow(query, productID).Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.Stock,
		&product.Category,
		&product.SKU,
		&imageURLsString, // Temporary variable for image_urls
		&product.Discount,
		&product.IsActive,
		&tagsString, // Temporary variable for tags
		&product.AdditionalInfo,
		&product.CreatedAt,
		&product.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("Product not found")
			return nil, fmt.Errorf("product not found")
		}
		log.Printf("Error querying product: %v", err)
		return nil, err
	}

	// Convert string fields to slices
	product.StringToImageURLs(imageURLsString)
	product.StringToTags(tagsString)

	// Return
	return &product, nil
}

// Update implements repository.ProductRepository.
func (r *ProductRepositoryImpl) Update(product *entity.Product) error {
	result, err := r.db.Exec(`
	UPDATE products
	SET name = $1, description = $2, price = $3, stock = $4, category = $5, sku = $6, image_urls = $7, discount = $8, is_active = $9, tags = $10, additional_info = $11, updated_at = $12
	WHERE id = $13`, product.Name, product.Description, product.Price, product.Stock, product.Category, product.SKU, product.ImageURLsToString(), product.Discount, product.IsActive, product.TagsToString(), product.AdditionalInfo, product.UpdatedAt, product.ID)

	if err != nil {
		log.Printf(" Error updating user: %v", err)
		return err
	}

	rowsAffeceted, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error getting rows affected: %v", err)
		return nil
	}

	if rowsAffeceted == 0 {
		log.Printf(" No rows affected")
		return nil
	}

	log.Printf(" Rows affected: %d", rowsAffeceted)
	return nil

}

// NewProductRepositoryImpl returns a new instance of ProductRepositoryImpl
func NewProductRepositoryImpl(db *sql.DB) repository.ProductRepository {
	return &ProductRepositoryImpl{db: db}
}
