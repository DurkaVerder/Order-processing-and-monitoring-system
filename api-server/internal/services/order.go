// This package contains the implementation of the order service.
package services

import (
	"api-server/internal/kafka"
	"api-server/internal/models"
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
	if err := s.producer.SendMessageForCreateOrder(kafka.TopicNewOrders, order, kafka.MaxRetries); err != nil {
		return err
	}
	return nil
}

// ChangeStatusOrder send msg in Kafka
func (s *ServiceManager) ChangeStatusOrder(id int, status string) error {
	order := models.StatusOrder{ID: id, Status: status}
	if err := s.producer.SendMessageForChangeStatusOrder(kafka.TopicOrderStatus, order, kafka.MaxRetries); err != nil {
		return err
	}
	return nil
}

func (s *ServiceManager) IsOrderStatusValid(status string) bool {
	if status == "created" || status == "processing" || status == "done" {
		return true
	}
	return false
}
