package routes

import (
	"Marketplace-backend/internal/interface_adopter/controller"
	"Marketplace-backend/internal/repository"
	"Marketplace-backend/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine, userController *controller.UserController, tokenRepo repository.TokenRepository) {
	// applly middleware to protect certain routes
	authMiddleware := middleware.AuthMiddleware(tokenRepo)

	// user-releted routes

	userGroup := router.Group("/users")
	{
		// pupblic routes
		userGroup.POST("", userController.RegisterUser)
		userGroup.POST("/authenticate", userController.AuthenticateUser)

		// protected routes
		userGroup.Use(authMiddleware)
		{
			userGroup.GET("", userController.GetAllUsers)
			userGroup.GET("/:id", userController.GetUserById)
			userGroup.PUT("/:id", userController.UpdateUser)
			userGroup.DELETE("/:id", userController.DeleteUser)
		}

	}

}
