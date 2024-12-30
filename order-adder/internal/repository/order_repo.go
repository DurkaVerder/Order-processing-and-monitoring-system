// This package contains the implementation of the AddOrder method of the repository interface.
package repository

import "Order-processing-and-monitoring-system/common/models"

// AddOrder adds an order to the repository and return id added order.
func (r *RepositoryManager) AddOrder(order models.Order) (int, error) {
	var orderID int
	req := "INSERT INTO orders (customer_name, customer_email, description, created_at, updated_at, amount) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	err := r.db.QueryRow(req, order.CustomerName, order.CustomerEmail, order.Description, order.CreatedAt, order.UpdateAt, order.Amount).Scan(&orderID)
	if err != nil {
		return 0, err
	}
	return orderID, nil
}
