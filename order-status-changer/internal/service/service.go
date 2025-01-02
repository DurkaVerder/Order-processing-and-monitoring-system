// This package contains the service interface and the service manager struct.
package service

import (
	"order-status-changer/internal/kafka/producer"
	"order-status-changer/internal/models"
	"order-status-changer/internal/repository"
)

// Service is an interface that defines the methods of the service.
type Service interface {
	AddStatusOrder(order models.StatusOrder) error
	ChangeStatusOrder(order models.StatusOrder) error
}

// ServiceManager is a struct that contains the Kafka and repository instances.
type ServiceManager struct {
	repo     repository.Repository
	producer producer.Producer
}

// NewServiceManager creates a new ServiceManager instance.
func NewServiceManager(repo repository.Repository, producer producer.Producer) *ServiceManager {
	return &ServiceManager{repo: repo, producer: producer}
}
