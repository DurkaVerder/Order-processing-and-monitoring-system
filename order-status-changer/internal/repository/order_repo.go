// This package contains the implementation of the AddOrder method of the repository interface.
package repository

import "Order-processing-and-monitoring-system/common/models"

// AddOrder adds an order to the repository and return id added order.
func (r *RepositoryManager) AddStatusOrder(order models.StatusOrder) error {
	req := "INSERT INTO status_order (order_id, status) VALUES ($1, $2)"
	_, err := r.db.Exec(req, order.ID, order.Status)
	if err != nil {
		return err
	}
	return nil
}

// ChangeStatusOrder changes the status of the order in the repository.
func (r *RepositoryManager) ChangeStatusOrder(order models.StatusOrder) error {
	req := "UPDATE status_order SET status = $1 WHERE order_id = $2"
	_, err := r.db.Exec(req, order.Status, order.ID)
	if err != nil {
		return err
	}
	return nil
}
