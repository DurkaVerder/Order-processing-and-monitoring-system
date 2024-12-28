// This package contains the server initialization and start functions.
package server

import (
	"api-server/internal/handlers"

	"github.com/gin-gonic/gin"
)

func InitAndStartServer(h handlers.Handlers) {
	r := gin.Default()
	r.GET("/orders", h.GetOrders)
	r.GET("/order/:id", h.GetStatusOrder)
	r.POST("/order", h.CreateOrder)
	r.PUT("/order/:id", h.ChangeStatusOrder)
	r.Run(":2222")
}
