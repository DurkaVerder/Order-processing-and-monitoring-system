// Package repository provides functions for working with the database
package repository

import (
	"Order-processing-and-monitoring-system/common/models"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// Interface Repository describes functions for working with the database
type Repository interface {
	GetOrders() ([]models.Order, error)
	GetStatusOrder(id int) (models.Order, error)
	CreateOrder(order models.Order) (models.Order, error)
}

// RepositoryManager is a struct that implements the Repository interface
type RepositoryManager struct {
	db *sql.DB
}

// NewRepositoryManager creates a new RepositoryManager
func NewRepositoryManager() *RepositoryManager {
	db, err := initDb()
	if err != nil {
		log.Fatal("Error in InitDb:", err.Error())
	}
	return &RepositoryManager{db: db}
}


func initDb() (*sql.DB, error) {
	db, err := sql.Open("postgres", "user=postgres password=durka dbname=Order-processing-and-monitoring-system port=5432 sslmode=disable")
	if err != nil {
		return db, err
	}
	if err = db.Ping(); err != nil {
		return db, err
	}
	return db, nil
}
