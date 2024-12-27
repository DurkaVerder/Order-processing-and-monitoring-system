package controller

import (
	"Order-processing-and-monitoring-system/internal/service"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	CreateOrder(c *gin.Context)
	GetStatusOrder(c *gin.Context)
	GetOrders(c *gin.Context)
	ChangeStatusOrder(c *gin.Context)
}

type ControllerManager struct {
	s *service.Service
}

func NewControllerManager(s *service.Service) *ControllerManager {
	return &ControllerManager{s: s}
}
