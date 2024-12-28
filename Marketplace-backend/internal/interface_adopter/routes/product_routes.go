package routes

import (
	"Marketplace-backend/internal/interface_adopter/controller"
	"Marketplace-backend/internal/repository"
	"Marketplace-backend/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(router *gin.Engine, productController controller.ProductController, tokeRepo repository.TokenRepository) {
	// apply middlware
	authMiddleware := middleware.AuthMiddleware(tokeRepo)

	// product-releted routes
	productRoutes := router.Group("/products")

	{
		// protected
		productRoutes.Use(authMiddleware)
		{
			// create product
			productRoutes.POST("/", productController.CreateProduct)
			// get all products
			productRoutes.GET("/", productController.GetAllProducts)
			// get product by id
			productRoutes.GET("/:id", productController.GetProductById)
			// update product
			productRoutes.PUT("/:id", productController.UpdateProduct)
			// delete product
			productRoutes.DELETE("/:id", productController.DeleteUser)

		}
	}
}
