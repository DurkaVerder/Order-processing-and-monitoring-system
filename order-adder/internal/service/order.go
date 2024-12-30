// This package contains the implementation of the AddOrder method of the service interface.
package service

import "Order-processing-and-monitoring-system/common/models"

// AddOrder adds an order to the repository.
func (s *ServiceManager) AddOrder(order models.Order) error {
	id, err := s.repo.AddOrder(order)
	if err != nil {
		return err
	}
	orderStatus := models.StatusOrder{ID: id, Status: "created"}
	if err = s.producer.SendMessageForSetStatusOrder("order.status", orderStatus); err != nil {
		return err
	}

	return nil
}
