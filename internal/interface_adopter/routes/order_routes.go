package routes

import (
	"Marketplace-backend/internal/interface_adopter/controller"
	"Marketplace-backend/internal/repository"
	"Marketplace-backend/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterOrderRoutes(router *gin.Engine, orderController controller.OrderController, tokenRepo repository.TokenRepository) {
	// apply auth
	authMiddleware := middleware.AuthMiddleware(tokenRepo)

	// order-releted routes
	orderRoutes := router.Group("/orders")
	{
		//protected
		orderRoutes.Use(authMiddleware)
		{
			orderRoutes.POST("/", orderController.CreateOrder)
			orderRoutes.GET("/", orderController.GetAllOrders)
			orderRoutes.GET("/:id", orderController.GetOrderByID)
			orderRoutes.PUT("/:id", orderController.UpdateOrder)
			orderRoutes.DELETE("/:id", orderController.DeleteOrder)

		}
	}
}
