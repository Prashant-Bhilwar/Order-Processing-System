package main

import (
	"github.com/prashant-bhilwar/order-processing-system/product-service/config"
	"github.com/prashant-bhilwar/order-processing-system/product-service/database"

	"github.com/prashant-bhilwar/order-processing-system/product-service/routes"

	"github.com/gin-gonic/gin"

	_ "github.com/prashant-bhilwar/order-processing-system/product-service/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Product Service API
// @version 1.0
// @description REST APIs for managing products
// @host localhost:8081
// @BasePath /
func main() {
	config.LoadConfig()
	database.InitPostgres()

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routes.RegisterRoutes(r)
	r.Run(":8081")
}
