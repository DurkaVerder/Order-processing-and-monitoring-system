// This package contains the repository interface and the repository manager.
package repository

import (
	"database/sql"
	"fmt"
	"log"
	"order-adder/config"
	"order-adder/internal/models"

	_ "github.com/lib/pq"
)

// Repository is the interface that wraps the basic methods for the repository.
type Repository interface {
	AddOrder(order models.Order) (int, error)
}

// RepositoryManager is the manager for the repository.
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
	connect := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.Database.Host, cfg.Database.User, cfg.Database.Password, cfg.Database.DBName, cfg.Database.Port, cfg.Database.SSLMode)

	db, err := sql.Open("postgres", connect)
	if err != nil {
		return db, err
	}
	if err = db.Ping(); err != nil {
		return db, err
	}
	return db, nil
}
