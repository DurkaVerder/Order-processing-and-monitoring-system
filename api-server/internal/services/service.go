// Package services provides functions for the operation of the service
package services

import (
	"Order-processing-and-monitoring-system/common/models"
	"api-server/internal/repository"
)

// Interface Service describes functions the operation of the service
type Service interface {
	GetOrders() ([]models.Order, error)
	GetStatusOrder(id int) (models.Order, error)
	CreateOrder(order models.Order) (models.Order, error)
	ChangeStatusOrder(id int, status string) (models.Order, error)
}

// ServiceManager is a struct that implements the Service interface
type ServiceManager struct {
	repo repository.Repository
}

// NewServiceManager creates a new ServiceManager
func NewServiceManager(repo repository.Repository) *ServiceManager {
	return &ServiceManager{repo: repo}
}
