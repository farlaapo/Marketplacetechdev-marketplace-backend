package controller

import (
	"Marketplace-backend/internal/entity"
	"Marketplace-backend/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

// ReviewRatingController represents the review rating controller
type ReviewRatingController struct {
	ReviewRatingService service.ReviewRatingService
}

// NewReviewRatingController will initialize the review rating controller
func NewReviewRatingController(ReviewRatingService service.ReviewRatingService) *ReviewRatingController {
	return &ReviewRatingController{
		ReviewRatingService: ReviewRatingService,
	}
}

func (Rc *ReviewRatingController) CreateReviewRating(c *gin.Context) {
	// Initialize review rating struct
	var reviewRating entity.ReviewRating

	// Bind the request body to the review rating struct
	if err := c.BindJSON(&reviewRating); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload", "details": err.Error()})
		return
	}
	// Call the service method with the correct order and all required arguments
	createdReviewRating, err := Rc.ReviewRatingService.CreateReviewRating(reviewRating.UserID, reviewRating.ProductID, reviewRating.Rating, reviewRating.Comment)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create review rating", "details": err.Error()})
		return
	}
	// Respond with the created review rating
	c.JSON(201, gin.H{"message": "Review rating created successfully", "reviewRating": createdReviewRating})
}

// UpdateReviewRating updates a review rating
func (Rc *ReviewRatingController) UpdateReviewRating(c *gin.Context) {
	// Initialize review rating struct
	var reviewRating entity.ReviewRating
	// parse the id from the request
	reviewRatingParam := c.Param("id")
	reviewRatingID, err := uuid.FromString(reviewRatingParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid review rating ID", "details": err.Error()})
		return
	}

	// Bind the request body to the review rating struct
	if err := c.BindJSON(&reviewRating); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload", "details": err.Error()})
		return
	}

	reviewRating.ID = reviewRatingID

	// call service to update the review rating
	if err := Rc.ReviewRatingService.UpdateReviewRating(&reviewRating); err != nil {
		c.JSON(500, gin.H{"error": "Failed to update review rating", "details": err.Error()})
		return
	}
	// Respond with the updated review rating
	c.JSON(200, gin.H{"message": "Review rating updated successfully", "reviewRating": reviewRating})
}

// DeleteReviewRating deletes a review rating
func (Rc *ReviewRatingController) DeleteReviewRating(c *gin.Context) {
	// parse the id from the request
	reviewRatingParam := c.Param("id")
	reviewRatingID, err := uuid.FromString(reviewRatingParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid review rating ID", "details": err.Error()})
		return
	}

	// call service to delete the review rating
	if err := Rc.ReviewRatingService.DeleteReviewRating(reviewRatingID); err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete review rating", "details": err.Error()})
		return
	}

	// Respond with the deleted review rating
	c.JSON(200, gin.H{"message": "Review rating deleted successfully"})
}

// GetByID gets a review rating by ID
func (Rc *ReviewRatingController) GetReviewRatingByID(c *gin.Context) {
	// parse the id from the request
	reviewRatingParam := c.Param("id")
	reviewRatingID, err := uuid.FromString(reviewRatingParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid review rating ID", "details": err.Error()})
		return
	}

	// call service to get the review rating
	reviewRating, err := Rc.ReviewRatingService.GetReviewRatingByID(reviewRatingID)
	if err != nil {
		c.JSON(404, gin.H{"error": "Review rating not found", "details": err.Error()})
		return
	}

	// Respond with the review rating
	c.JSON(200, gin.H{"reviewRating": reviewRating})

}

// GetReviewRatings gets all review ratings
func (Rc *ReviewRatingController) GetReviewRatings(c *gin.Context) {
	// call service to get all review ratings
	reviewRatings, err := Rc.ReviewRatingService.GetReviewRatings()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get review ratings", "details": err.Error()})
		return
	}
	// Respond with the review ratings
	c.JSON(200, gin.H{"reviewRatings": reviewRatings})
}
