// This package contains the implementation of the AddOrder method of the service interface.
package service

import (
	"order-status-changer/internal/models"
	"order-status-changer/internal/kafka"
	"time"
)

func (s *ServiceManager) AddStatusOrder(order models.StatusOrder) error {
	err := s.repo.AddStatusOrder(order)
	if err != nil {
		return err
	}
	report := models.Report{Status: order.Status, DateTime: time.Now()}
	if err := s.producer.SendMessageForAnalytics(kafka.TopicReport, report, kafka.MaxRetries); err != nil {
		return err
	}
	return nil
}

// AddOrder adds an order to the repository.
func (s *ServiceManager) ChangeStatusOrder(order models.StatusOrder) error {
	err := s.repo.ChangeStatusOrder(order)
	if err != nil {
		return err
	}
	err = s.repo.ChangeUpdateDateStatus(order.ID)
	if err != nil {
		return err
	}

	report := models.Report{Status: order.Status, DateTime: time.Now()}
	if err := s.producer.SendMessageForAnalytics(kafka.TopicReport, report, kafka.MaxRetries); err != nil {
		return err
	}

	return nil
}
