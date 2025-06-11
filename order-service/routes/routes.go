package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/prashant-bhilwar/order-processing-system/order-service/handler"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/orders", handler.CreateOrder)
	r.GET("/orders/:userId", handler.GetOrders)
}
