package routes

import (
	"Marketplace-backend/internal/interface_adopter/controller"

	"Marketplace-backend/internal/repository"
	"Marketplace-backend/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterSalesRoutes(router *gin.Engine, salesController controller.SalesController, tokenRepo repository.TokenRepository) {
	//aply mid
	authMiddleware := middleware.AuthMiddleware(tokenRepo)

	// sales-releted
	salesRouter := router.Group("/sales")
	{
		// protecetd
		salesRouter.Use(authMiddleware)
		{
			// Create Sales
			salesRouter.POST("/", salesController.CreateSales)
			// Update Sales
			salesRouter.PUT("/:id", salesController.UpdateSales)
			// Get Sales
			salesRouter.GET("/", salesController.GetAllSales)
			// Get Sales by ID
			salesRouter.GET("/:id", salesController.GetSalesByID)
			// Delete Sales
			salesRouter.DELETE("/:id", salesController.DeleteSales)
		}
	}

}
