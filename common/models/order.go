package models

type Order struct {
	ID            int     `json:"id"`
	CustomerName  string  `json:"customer_name"`
	CustomerEmail string  `json:"customer_email"`
	Description   string  `json:"description"`
	Status        string  `json:"status"`
	CreatedAt     string  `json:"created_at"`
	UpdateAt      string  `json:"update_at"`
	Amount        float64 `json:"amount"`
}

type StatusOrder struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}

