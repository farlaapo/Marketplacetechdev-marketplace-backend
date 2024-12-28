package controller

import (
	"Marketplace-backend/internal/entity"
	"Marketplace-backend/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

// Controller is a struct that holds the service instance
type ProductController struct {
	productService service.ProductService
}

// NewProductController returns a new instance of ProductController
func NewProductController(productService service.ProductService) *ProductController {
	return &ProductController{productService: productService}
}

func (Pc *ProductController) CreateProduct(c *gin.Context) {
	// Initialize product struct
	var product entity.Product

	// Bind JSON payload to product struct
	if err := c.BindJSON(&product); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload", "details": err.Error()})
		return
	}

	// Call the service method with the correct order and all required arguments
	createdProduct, err := Pc.productService.CreateProduct(
		product.Name,          // string
		product.Description,   // string
		product.Price,         // float64
		product.Stock,         // int
		product.Category,      // string
		product.SKU,           // string
		product.ImageURLs,     // []string
		product.Discount,      // float64
		product.IsActive,      // bool
		product.Tags,          // []string
		product.AdditionalInfo, // string
	)

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create product", "details": err.Error()})
		return
	}

	// Respond with the created product
	c.JSON(201, gin.H{"message": "Product created successfully", "product": createdProduct})
}

func (Pc *ProductController) GetAllProducts(c *gin.Context) {
	// call the service
	product, err := Pc.productService.GetAllProducts()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	// return sucessfull
	c.JSON(200, gin.H{"products": product})
}

func (Pc *ProductController) GetProductById(c *gin.Context) {
	// Get the product ID from the URL parameter
	productParam := c.Param("id")
	productID, err := uuid.FromString(productParam)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	//  call service
	product, err := Pc.productService.GetProductById(productID)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	// return the product
	c.JSON(200, gin.H{"product": product})

}

func (Pc *ProductController) UpdateProduct(c *gin.Context) {
	var product entity.Product
	// Get the product ID from the URL parameter
	productParam := c.Param("id")
	productID, err := uuid.FromString(productParam)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// bind
	if err := c.BindJSON(&product); err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	product.ID = productID

	// call service
	if err := Pc.productService.UpdateProduct(product); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// return
	c.JSON(200, product)

}

func (Pc *ProductController) DeleteUser(c *gin.Context) {
	// Get the product ID from the URL parameter
	productParam := c.Param("id")
	productID, err := uuid.FromString(productParam)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// call service
	if err := Pc.productService.DeleteProduct(productID); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// return
	c.JSON(200, gin.H{"message": "product deleted"})

}
