package repository

import (
	"Marketplace-backend/internal/entity"

	"github.com/gofrs/uuid"
)

type ReviewRatingRepository interface {
	Create(reviewRating *entity.ReviewRating) error
	Update(reviewRating *entity.ReviewRating) error
	Delete(reviewRatingID uuid.UUID) error
	GetByID(reviewRatingID uuid.UUID) (*entity.ReviewRating, error)
	ListAll() ([]*entity.ReviewRating, error)
}
