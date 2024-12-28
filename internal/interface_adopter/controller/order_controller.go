package controller

import (
	"Marketplace-backend/internal/entity"
	"Marketplace-backend/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

// struct for order controller
type OrderController struct {
	orderService service.OrderService
}

// NewOrderController returns a new instance of OrderController
func NewOrderController(orderService service.OrderService) *OrderController {
	return &OrderController{orderService: orderService}
}

// CreateOrder creates a new order
func (Oc *OrderController) CreateOrder(c *gin.Context) {
	var order entity.Order
	// Bind JSON payload to order struct
	if err := c.BindJSON(&order); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// Call service to create order
	createdOrder, err := Oc.orderService.CreateOrder(order.UserID, order.ProductID, order.Quantity, order.Status)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Respond with the created order
	c.JSON(201, gin.H{"message": "Order created successfully", "order": createdOrder})

}

// GetAllOrders gets all orders
func (Oc *OrderController) GetAllOrders(c *gin.Context) {
	// Call
	order, err := Oc.orderService.GetAllOrders()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// respond
	c.JSON(200, order)

}

// GETBYD ORDER
func (Oc *OrderController) GetOrderByID(c *gin.Context) {
	orderParam := c.Param("id")
	orderID, err := uuid.FromString(orderParam)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// call service
	orderCreated, err := Oc.orderService.GetOrderByID(orderID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// respond
	c.JSON(200, orderCreated)

}

// Update order
func (Oc *OrderController) UpdateOrder(c *gin.Context) {
	var order entity.Order

	orderParam := c.Param("id")
	orderID, err := uuid.FromString(orderParam)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// bind json
	if err := c.BindJSON(&order); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	order.ID = orderID

	// call service
	if err := Oc.orderService.UpdateOrder(&order); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// respond
	c.JSON(200, gin.H{"error": "Order updated successfully"})

}

// Delete order
func (Oc *OrderController) DeleteOrder(c *gin.Context) {
	// param
	orderParam := c.Param("id")
	orderID, err := uuid.FromString(orderParam)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	// call service
	if err := Oc.orderService.DeleteOrder(orderID); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	// respond
	c.JSON(200, gin.H{"error": "Order deleted successfully"})
}
