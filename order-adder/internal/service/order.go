// This package contains the implementation of the AddOrder method of the service interface.
package service

import (
	"Order-processing-and-monitoring-system/common/models"
	"order-adder/internal/kafka"
	"time"
)

// AddOrder adds an order to the repository.
func (s *ServiceManager) AddOrder(order models.Order) error {
	order.CreatedAt = time.Now()
	id, err := s.repo.AddOrder(order)
	if err != nil {
		return err
	}
	orderStatus := models.StatusOrder{ID: id, Status: "created"}
	if err = s.producer.SendMessageForSetStatusOrder(kafka.TopicOrderStatus, orderStatus, kafka.MaxRetries); err != nil {
		return err
	}

	return nil
}
