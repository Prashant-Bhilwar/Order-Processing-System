package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/prashant-bhilwar/order-processing-system/api-gateway/handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.Any("/user/*proxyPath", handlers.UserProxy)
		api.Any("/product/*proxyPath", handlers.ProductProxy)
		api.Any("/order/*proxyPath", handlers.OrderProxy)
		api.Any("/cart/*proxyPath", handlers.CartProxy)
	}

	return r
}
