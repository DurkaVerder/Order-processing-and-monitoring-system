// This package contains the implementation of the order service.
package services

import (
	"Order-processing-and-monitoring-system/common/models"
	"time"
)

func (s *ServiceManager) GetOrders() ([]models.Order, error) {
	orders, err := s.repo.GetOrders()
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (s *ServiceManager) GetStatusOrder(id int) (string, error) {
	order, err := s.repo.GetOrder(id)
	if err != nil {
		return "", err
	}
	return order.Status, nil
}

func (s *ServiceManager) CreateOrder(order models.Order) error {
	order.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	order.UpdateAt = time.Now().Format("2006-01-02 15:04:05")
	order.Status = "Created"
	err := s.repo.CreateOrder(order)
	if err != nil {
		return err
	}
	return nil
}

func (s *ServiceManager) ChangeStatusOrder(id int, status string) error {
	return nil
}
