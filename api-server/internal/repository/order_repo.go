// This package contains the implementation of the order repository interface. 
// It contains the methods to interact with the database for the order entity.
package repository

import "Order-processing-and-monitoring-system/common/models"

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
		err := rows.Scan(&order.ID, &order.CustomerName, &order.CustomerEmail, &order.Description, &order.Status, &order.CreatedAt, &order.UpdateAt, &order.Amount)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

func (r *RepositoryManager) GetOrder(id int) (models.Order, error) {
	req := "SELECT * FROM orders WHERE id = $1"
	row := r.db.QueryRow(req, id)

	order := models.Order{}
	err := row.Scan(&order.ID, &order.CustomerName, &order.CustomerEmail, &order.Description, &order.Status, &order.CreatedAt, &order.UpdateAt, &order.Amount)
	if err != nil {
		return models.Order{}, err
	}

	return order, nil
}

func (r *RepositoryManager) CreateOrder(order models.Order) error {
	req := "INSERT INTO orders (customer_name, customer_email, description, status, created_at, update_at, amount) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	_, err := r.db.Exec(req, order.CustomerName, order.CustomerEmail, order.Description, order.Status, order.CreatedAt, order.UpdateAt, order.Amount)
	if err != nil {
		return err
	}

	return nil
}
