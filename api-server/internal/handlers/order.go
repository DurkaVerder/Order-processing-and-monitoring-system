// This package contains the implementation of the order handlers.
package handlers

import (
	"api-server/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Get all orders
func (h *HandlersManager) GetOrders(c *gin.Context) {
	orders, err := h.s.GetOrders()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, orders)
}

// Get the order status
func (h *HandlersManager) GetStatusOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	status, err := h.s.GetStatusOrder(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(200, gin.H{"status": status})
}

// Create an order
func (h *HandlersManager) CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.s.CreateOrder(order); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"info": "Order created"})
}

// Change the order status
func (h *HandlersManager) ChangeStatusOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	var request struct {
		Status string `json:"status" binding:"required"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if h.s.IsOrderStatusValid(request.Status) == false {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status"})
		return
	}

	if err := h.s.ChangeStatusOrder(id, request.Status); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"info": "Status changed"})
}
