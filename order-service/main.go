package main

import (
	_ "github.com/prashant-bhilwar/order-processing-system/order-service/docs"

	"github.com/gin-gonic/gin"
	"github.com/prashant-bhilwar/order-processing-system/order-service/config"
	"github.com/prashant-bhilwar/order-processing-system/order-service/database"
	"github.com/prashant-bhilwar/order-processing-system/order-service/routes"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Order Service API
// @version 1.0
// @description REST API for order placement and retrieval
// @host localhost:8084
// @BasePath /
func main() {
	config.LoadConfig()
	database.InitPostgres()

	r := gin.Default()
	routes.SetupRoutes(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8084")
}
