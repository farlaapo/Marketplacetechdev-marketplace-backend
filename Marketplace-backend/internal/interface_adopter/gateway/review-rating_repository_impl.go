package gateway

import (
	"Marketplace-backend/internal/entity"
	"Marketplace-backend/internal/repository"
	"database/sql"
	"fmt"
	"log"

	"github.com/gofrs/uuid"
)

// ReviewRatingRepositoryImpl represents the token repository contract
type ReviewRatingRepositoryImpl struct {
	db *sql.DB
}

// Create implements repository.ReviewRatingRepository.
func (r *ReviewRatingRepositoryImpl) Create(reviewRating *entity.ReviewRating) error {
	// generete UUID

	newUUD, err := uuid.NewV4()
	if err != nil {
		log.Printf(" ERROR uuid : %v", err)
		return err
	}
	reviewRating.ID = newUUD
	query := `INSERT INTO review_rating (id, user_id, product_id, rating, comment, created_at)
	VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;`
	err = r.db.QueryRow(query, reviewRating.ID, reviewRating.UserID, reviewRating.ProductID, reviewRating.Rating, reviewRating.Comment, reviewRating.CreatedAt).Scan(&reviewRating.ID)
	if err != nil {
		log.Printf("Error creating review rating: %v", err)
		return err
	}
	return nil
}

// Delete implements repository.ReviewRatingRepository.
func (r *ReviewRatingRepositoryImpl) Delete(reviewRatingID uuid.UUID) error {
	query := "DELETE FROM review_rating WHERE id = $1"
	result, err := r.db.Exec(query, reviewRatingID)
	if err != nil {
		log.Printf(" Error deleting review rating: %v", err)
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
	return nil
}

// GetByID implements repository.ReviewRatingRepository.
func (r *ReviewRatingRepositoryImpl) GetByID(reviewRatingID uuid.UUID) (*entity.ReviewRating, error) {
	// define the struct
	var reviewRating entity.ReviewRating
	query := "SELECT id, user_id, product_id, rating, comment, created_at, updated_at FROM review_rating WHERE id = $1"
	err := r.db.QueryRow(query, reviewRatingID).Scan(&reviewRating.ID, &reviewRating.UserID, &reviewRating.ProductID, &reviewRating.Rating, &reviewRating.Comment, &reviewRating.CreatedAt, &reviewRating.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf(" No rows found")
			return nil, fmt.Errorf("no rows found")
		}
		log.Printf("Error getting review rating: %v", err)
		return nil, err
	}
	return &reviewRating, nil

}

// ListAll implements repository.ReviewRatingRepository.
func (r *ReviewRatingRepositoryImpl) ListAll() ([]*entity.ReviewRating, error) {
	rows, err := r.db.Query("SELECT id, user_id, product_id, rating, comment, created_at, updated_at FROM review_rating")
	if err != nil {
		log.Printf("Error getting review rating: %v", err)
		return nil, err
	}

	defer rows.Close()
	var reviewRatings []*entity.ReviewRating
	for rows.Next() {
		var reviewRating entity.ReviewRating
		err := rows.Scan(&reviewRating.ID, &reviewRating.UserID, &reviewRating.ProductID, &reviewRating.Rating, &reviewRating.Comment, &reviewRating.CreatedAt, &reviewRating.UpdatedAt)
		if err != nil {
			log.Printf("Error scanning review rating: %v", err)
			return nil, err
		}
		reviewRatings = append(reviewRatings, &reviewRating)

	}
	return reviewRatings, nil
}

// Update implements repository.ReviewRatingRepository.
func (r *ReviewRatingRepositoryImpl) Update(reviewRating *entity.ReviewRating) error {
	query := `UPDATE review_rating SET user_id = $1, product_id = $2, rating = $3, comment = $4, updated_at = $5 WHERE id = $6`
	result, err := r.db.Exec(query, reviewRating.UserID, reviewRating.ProductID, reviewRating.Rating, reviewRating.Comment, reviewRating.UpdatedAt, reviewRating.ID)
	if err != nil {
		log.Printf("Error updating review rating: %v", err)
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
	return nil
}

// NewReviewRatingRepositoryImpl creates a new instance of ReviewRatingRepositoryImpl
func NewReviewRatingRepositoryImpl(db *sql.DB) repository.ReviewRatingRepository {
	return &ReviewRatingRepositoryImpl{db: db}
}
