package models

import "time"

type Order struct {
	ID            int       `json:"id"`
	CustomerName  string    `json:"customer_name"`
	CustomerEmail string    `json:"customer_email"`
	Description   string    `json:"description"`
	CreatedAt     time.Time `json:"created_at"`
	UpdateAt      time.Time `json:"update_at"`
	Amount        float64   `json:"amount"`
}

type StatusOrder struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}

type Report struct {
	Status   string    `json:"status"`
	DateTime time.Time `json:"date_time"`
}
