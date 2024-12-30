// This package contains the implementation of the AddOrder method of the service interface.
package service

import (
	"Order-processing-and-monitoring-system/common/models"
	"time"
)

// AddOrder adds an order to the repository.
func (s *ServiceManager) AddOrder(order models.Order) error {
	order.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	id, err := s.repo.AddOrder(order)
	if err != nil {
		return err
	}
	orderStatus := models.StatusOrder{ID: id, Status: "created"}
	if err = s.producer.SendMessageForSetStatusOrder("orders.status", orderStatus); err != nil {
		return err
	}

	return nil
}
