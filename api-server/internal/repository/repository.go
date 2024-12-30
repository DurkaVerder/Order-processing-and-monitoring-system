// This package contains the Repository interface and the RepositoryManager struct that implements it.
// The RepositoryManager struct is responsible for working with the database.
package repository

import (
	"Order-processing-and-monitoring-system/common/models"
	"api-server/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Interface Repository describes functions for working with the database
type Repository interface {
	GetOrders() ([]models.Order, error)
	GetOrderStatus(id int) (models.StatusOrder, error)
}

// RepositoryManager is a struct that implements the Repository interface
type RepositoryManager struct {
	db *sql.DB
}

// NewRepositoryManager creates a new RepositoryManager
func NewRepositoryManager(cfg *config.Config) *RepositoryManager {
	db, err := initDb(cfg)
	if err != nil {
		log.Fatal("Error in InitDb:", err.Error())
	}
	return &RepositoryManager{db: db}
}

func initDb(cfg *config.Config) (*sql.DB, error) {
	connect := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.Database.User, cfg.Database.Password, cfg.Database.DBName, cfg.Database.Port, cfg.Database.SSLMode)
	db, err := sql.Open("postgres", connect)
	if err != nil {
		return db, err
	}
	if err = db.Ping(); err != nil {
		return db, err
	}
	return db, nil
}
