package handlers

import (
	service "api-server/internal/services"

	"github.com/gin-gonic/gin"
)

type Handlers interface {
	GetOrders(c *gin.Context)
	GetStatusOrder(c *gin.Context)
	CreateOrder(c *gin.Context)
	ChangeStatusOrder(c *gin.Context)
}

type HandlersManager struct {
	s service.Service
}

func NewHandlersManager(s service.Service) *HandlersManager {
	return &HandlersManager{s}
}
