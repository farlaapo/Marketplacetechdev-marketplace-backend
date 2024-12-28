package controller

import (
	"Marketplace-backend/internal/entity"
	"Marketplace-backend/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

// Controller is a struct that holds the service instance
type SalesController struct {
	salesService service.SalesService
}

// New SalesController returns a new instance of SalesController
func NewSalesContriller(salesService service.SalesService) *SalesController {
	return &SalesController{
		salesService: salesService,
	}
}

// CreateSales
func (Sc *SalesController) CreateSales(c *gin.Context) {
	var sales entity.Sales
	// bind
	if err := c.BindJSON(&sales); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// call service
	createSales, err := Sc.salesService.CreateSales(sales.UserID, sales.OrdeID, sales.ProductID, sales.TotalSales, sales.TotalPrice, sales.TotalOrders, sales.Quantity, sales.OrderDate, sales.Status)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	// return
	c.JSON(201, createSales)
}

// GetSales returns a list of sales
func (Sc *SalesController) GetAllSales(c *gin.Context) {
	// call service
	sales, err := Sc.salesService.GetAllSales()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	// return
	c.JSON(200, sales)
}

// Get SALES BY id
func (Sc *SalesController) GetSalesByID(c *gin.Context) {
	// get id param
	salesParam := c.Param("id")
	salesID, err := uuid.FromString(salesParam)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// call service
	sales, err := Sc.salesService.GetSalesByID(salesID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// return
	c.JSON(200, sales)
}

// UPDATE Sales
func (Sc *SalesController) UpdateSales(c *gin.Context) {
	var sales entity.Sales
	// get id param
	salesParam := c.Param("id")
	salesID, err := uuid.FromString(salesParam)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// bind
	if err := c.BindJSON(&sales); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	sales.ID = salesID

	// call service
	if err := Sc.salesService.UpdateSales(sales); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// return
	c.JSON(200, gin.H{"massage": "Sales updated successfully"})

}

func (Sc *SalesController) DeleteSales(c *gin.Context) {
	// get id param

	salesParam := c.Param("id")
	salesID, err := uuid.FromString(salesParam)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// call service
	if err := Sc.salesService.DeleteSales(salesID); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// return
	c.JSON(200, gin.H{"massage": " Sales deleteed successfully"})
}
