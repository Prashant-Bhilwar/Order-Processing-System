package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/prashant-bhilwar/order-processing-system/product-service/handler"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/products", handler.CreateProduct)
	r.GET("/products", handler.GetAllProducts)
}
