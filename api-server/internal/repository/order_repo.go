// This package contains the implementation of the order repository interface.
// It contains the methods to interact with the database for the order entity.
package repository

import "api-server/internal/models"

// GetOrders returns all orders from the database.
func (r *RepositoryManager) GetOrders() ([]models.Order, error) {
	req := "SELECT * FROM orders"
	rows, err := r.db.Query(req)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := []models.Order{}
	for rows.Next() {
		order := models.Order{}
		err := rows.Scan(&order.ID, &order.CustomerName, &order.CustomerEmail, &order.Description, &order.CreatedAt, &order.UpdateAt, &order.Amount)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

// GetOrder returns the order by id from the database.
func (r *RepositoryManager) GetOrderStatus(id int) (models.StatusOrder, error) {
	req := "SELECT * FROM order-status WHERE id = $1"
	row := r.db.QueryRow(req, id)

	order := models.StatusOrder{}
	err := row.Scan(&order.ID, &order.Status)
	if err != nil {
		return models.StatusOrder{}, err
	}

	return order, nil
}
