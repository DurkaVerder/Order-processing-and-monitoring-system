// This package contains the implementation of the order service.
package services

import (
	"Order-processing-and-monitoring-system/common/models"
)

// GetOrders returns all orders.
func (s *ServiceManager) GetOrders() ([]models.Order, error) {
	orders, err := s.repo.GetOrders()
	if err != nil {
		return nil, err
	}
	return orders, nil
}

// GetStatusOrder returns the status of the order by id.
func (s *ServiceManager) GetStatusOrder(id int) (string, error) {
	order, err := s.repo.GetOrderStatus(id)
	if err != nil {
		return "", err
	}
	return order.Status, nil
}

// CreateOrder send msg in Kafka
func (s *ServiceManager) CreateOrder(order models.Order) error {
	if err := s.producer.SendMessageForCreateOrder("orders.new", order, 5); err != nil {
		return err
	}
	return nil
}

// ChangeStatusOrder send msg in Kafka
func (s *ServiceManager) ChangeStatusOrder(id int, status string) error {
	order := models.StatusOrder{ID: id, Status: status}
	if err := s.producer.SendMessageForChangeStatusOrder("orders.status", order, 5); err != nil {
		return err
	}
	return nil
}
