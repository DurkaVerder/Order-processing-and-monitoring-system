// This package is the entry point of the application. It will start the server and listen for incoming requests.
package main

import (
	"api-server/config"
	"api-server/internal/handlers"
	"api-server/internal/repository"
	"api-server/internal/server"
	"api-server/internal/services"
)

func main() {
	// Initialize the configuration
	cfg := config.InitConfig()

	// Create a new repository
	repository := repository.NewRepositoryManager(cfg)

	// Create a new service
	service := services.NewServiceManager(repository)

	// Create a new handler
	handler := handlers.NewHandlersManager(service)

	// Start the server
	server.InitAndStartServer(handler, cfg)
}
