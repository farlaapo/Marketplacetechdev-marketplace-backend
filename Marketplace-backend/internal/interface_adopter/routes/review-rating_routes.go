package routes

import (
	"Marketplace-backend/internal/interface_adopter/controller"
	"Marketplace-backend/internal/repository"
	"Marketplace-backend/pkg/middleware"

	"github.com/gin-gonic/gin"
)


func RegisterReviewRatingRoutes(router *gin.Engine, reviewRatingController controller.ReviewRatingController, tokenRepo repository.TokenRepository) {
	// apply middleware
	authMiddleware := middleware.AuthMiddleware(tokenRepo)

	// review-rating-related routes
	reviewRatingRoutes := router.Group("/review-rating")
	{
		// protected
		reviewRatingRoutes.Use(authMiddleware)
		{
			// create review-rating
			reviewRatingRoutes.POST("/", reviewRatingController.CreateReviewRating)
			// get all review-rating
			reviewRatingRoutes.GET("/", reviewRatingController.GetReviewRatings)
			// get review-rating by id
			reviewRatingRoutes.GET("/:id", reviewRatingController.GetReviewRatingByID)
			// update review-rating
			reviewRatingRoutes.PUT("/:id", reviewRatingController.UpdateReviewRating)
			// delete review-rating
			reviewRatingRoutes.DELETE("/:id", reviewRatingController.DeleteReviewRating)
		}
	}
}