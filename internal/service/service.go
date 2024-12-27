package service

import (
	"Order-processing-and-monitoring-system/internal/model"
	"Order-processing-and-monitoring-system/internal/storage"
)

type Service interface {
	CreateOrder(order model.Order) (int, error)
	GetStatusOrder(id int) (string, error)
	GetOrders() ([]model.Order, error)
	ChangeStatusOrder(id int, status string) error
}

type ServiceManager struct {
	storage storage.Storage
}

func NewServiceManager(storage storage.Storage) *ServiceManager {
	return &ServiceManager{storage: storage}
}
