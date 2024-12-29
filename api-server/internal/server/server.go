// This package contains the server initialization and start functions.
package server

import (
	"api-server/config"
	"api-server/internal/handlers"

	"github.com/gin-gonic/gin"
)

// InitAndStartServer initializes and starts the server.
func InitAndStartServer(h handlers.Handlers, cfg *config.Config) {
	r := gin.Default()
	r.GET("/orders", h.GetOrders)
	r.GET("/order/:id", h.GetStatusOrder)
	r.POST("/order", h.CreateOrder)
	r.PUT("/order/:id", h.ChangeStatusOrder)
	r.Run(cfg.Server.Port)
}
