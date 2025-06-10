package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prashant-bhilwar/order-processing-system/cart-service/config"
	"github.com/prashant-bhilwar/order-processing-system/cart-service/database"
	"github.com/prashant-bhilwar/order-processing-system/cart-service/routes"

	_ "github.com/prashant-bhilwar/order-processing-system/cart-service/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Cart Service API
// @version 1.0
// @description API for managing user cart operations
// @host localhost:8083
// @BasePath /
// @schemes http
func main() {
	config.LoadConfig()
	database.InitPostgres()

	r := gin.Default()
	routes.SetupRoutes(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8083")
}
