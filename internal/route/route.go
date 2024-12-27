package route

import (
	"Order-processing-and-monitoring-system/internal/controller"

	"github.com/gin-gonic/gin"
)

func RunServer(c controller.Controller) {
	r := gin.Default()
	r.GET("/orders", c.GetOrders)
	r.GET("/order/:id", c.GetStatusOrder)
	r.POST("/order", c.CreateOrder)
	r.PUT("/order/:id", c.ChangeStatusOrder)
	r.Run(":2222")
}
