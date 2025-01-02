// This package contains the Service interface and ServiceManager struct that implements the Service interface.
package services

import (
	"api-server/internal/kafka/producer"
	"api-server/internal/models"
	"api-server/internal/repository"
)

// Interface Service describes functions the operation of the service
type Service interface {
	GetOrders() ([]models.Order, error)
	GetStatusOrder(id int) (string, error)
	CreateOrder(order models.Order) error
	ChangeStatusOrder(id int, status string) error
	IsOrderStatusValid(status string) bool
}

// ServiceManager is a struct that implements the Service interface
type ServiceManager struct {
	repo     repository.Repository
	producer producer.Producer
}

// NewServiceManager creates a new ServiceManager
func NewServiceManager(producer producer.Producer, repo repository.Repository) *ServiceManager {
	return &ServiceManager{producer: producer, repo: repo}
}
