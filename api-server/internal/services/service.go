package services

import "Order-processing-and-monitoring-system/common/models"

type Service interface {
	GetOrders() ([]models.Order, error)
	GetStatusOrder(id int) (models.Order, error)
	CreateOrder(order models.Order) (models.Order, error)
	ChangeStatusOrder(id int, status string) (models.Order, error)
}
