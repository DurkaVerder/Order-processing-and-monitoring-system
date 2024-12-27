package storage

import (
	"Order-processing-and-monitoring-system/internal/model"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateOrder(order model.Order) (int, error)
	GetStatusOrder(id int) (string, error)
	GetOrders() ([]model.Order, error)
	ChangeStatusOrder(id int, status string) error
}

type StorageManager struct {
	db *sql.DB
}

func NewStorageManager() *StorageManager {
	db, err := sql.Open("postgres", "user=postgres password=durka dbname=Order-processing-and-monitoring-system port=5432 sslmode=disable")
	if err != nil {
		log.Fatal("Error open db: ", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Error connect db: ", err)
	}
	return &StorageManager{db: db}
}
