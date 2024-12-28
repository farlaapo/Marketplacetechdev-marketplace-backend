package service

import (
	"Marketplace-backend/internal/entity"
	"Marketplace-backend/internal/repository"
	"time"
	"github.com/gofrs/uuid"
)

// ReviewRatingService is a set of methods used to manipulate review
type ReviewRatingService interface {
	CreateReviewRating(UserID, ProductID uuid.UUID, Rating int, Comment string) (*entity.ReviewRating, error)
	UpdateReviewRating(reviewRating *entity.ReviewRating) error
	GetReviewRatingByID(reviewRatingID uuid.UUID) (*entity.ReviewRating, error)
	GetReviewRatings() ([]*entity.ReviewRating, error)
	DeleteReviewRating(reviewRatingID uuid.UUID) error
}

// reviewRatingService implements ReviewRatingService struct
type reviewRatingService struct {
	repo      repository.ReviewRatingRepository
	tokenRepo repository.TokenRepository
}

// CreateReviewRatin implements ReviewRatingService.
func (s *reviewRatingService) CreateReviewRating(UserID uuid.UUID, ProductID uuid.UUID, Rating int, Comment string) (*entity.ReviewRating, error) {
	// generate UUID
	newUUID, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	reviewRating := &entity.ReviewRating{
		ID:        newUUID,
		UserID:    UserID,
		ProductID: ProductID,
		Rating:    Rating,
		Comment:   Comment,
		CreatedAt: time.Now(),
	}

	if err := s.repo.Create(reviewRating); err != nil {
		return nil, err
	}	
	return reviewRating, nil
}

// DeleteReviewRating implements ReviewRatingService.
func (s *reviewRatingService) DeleteReviewRating(reviewRatingID uuid.UUID) error {
	_, err := s.repo.GetByID(reviewRatingID)
	if err != nil {
			return err
	}
	if err := s.repo.Delete(reviewRatingID); err != nil {
		return err
	}
	return nil
}

// GetReviewRating implements ReviewRatingService.
func (s *reviewRatingService) GetReviewRatingByID(reviewRatingID uuid.UUID) (*entity.ReviewRating, error) {
	reviewRating, err := s.repo.GetByID(reviewRatingID)
	if err != nil {
		return nil, err
	}
	return reviewRating, nil
}

// GetReviewRatings implements ReviewRatingService.
func (s *reviewRatingService) GetReviewRatings() ([]*entity.ReviewRating, error) {
	reviewRatings, err := s.repo.ListAll()
	if err != nil {
		return nil, err
	}
  return reviewRatings, nil
}

// UpdateReviewRating implements ReviewRatingService.
func (s *reviewRatingService) UpdateReviewRating(reviewRating *entity.ReviewRating) error {
	_, err := s.repo.GetByID(reviewRating.ID)
	if err != nil {
		return err
	}
	if err := s.repo.Update(reviewRating); err != nil {
		return err
	}
	return nil
}

// NewReviewRatingService creates a new review rating service
func NewReviewRatingService(reviewRatingRepo repository.ReviewRatingRepository, tokenRepo repository.TokenRepository) ReviewRatingService {
	return &reviewRatingService{
		repo:      reviewRatingRepo,
		tokenRepo: tokenRepo,
	}
}
