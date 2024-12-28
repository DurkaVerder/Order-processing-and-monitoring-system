// This package is the entry point of the application. It will start the server and listen for incoming requests.
package main

import (
	"api-server/internal/handlers"
	"api-server/internal/repository"
	"api-server/internal/server"
	"api-server/internal/services"
)

func main() {
	// Create a new repository
	repository := repository.NewRepositoryManager()

	service := services.NewServiceManager(repository)
	// Create a new handler
	handler := handlers.NewHandlersManager(service)

	// Start the server
	server.InitAndStartServer(handler)
}
