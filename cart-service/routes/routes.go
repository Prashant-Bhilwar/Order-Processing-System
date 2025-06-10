package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/prashant-bhilwar/order-processing-system/cart-service/handler"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/cart", handler.AddToCart)
	r.DELETE("/cart/:userId/:productId", handler.RemoveFromCart)
	r.GET("/cart/:userId", handler.GetCart)
}
