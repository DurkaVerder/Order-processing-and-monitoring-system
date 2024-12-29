// This package contains the handlers interface and the handlers manager struct.
package handlers

import (
	service "api-server/internal/services"

	"github.com/gin-gonic/gin"
)

// Handlers interface describes the methods of the handlers.
type Handlers interface {
	GetOrders(c *gin.Context)
	GetStatusOrder(c *gin.Context)
	CreateOrder(c *gin.Context)
	ChangeStatusOrder(c *gin.Context)
}

// HandlersManager is a struct that implements the Handlers interface.
type HandlersManager struct {
	s service.Service
}


// NewHandlersManager creates a new HandlersManager.
func NewHandlersManager(s service.Service) *HandlersManager {
	return &HandlersManager{s}
}


