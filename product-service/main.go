package main

import (
	"github.com/prashant-bhilwar/order-processing-system/product-service/config"
	"github.com/prashant-bhilwar/order-processing-system/product-service/database"

	"github.com/prashant-bhilwar/order-processing-system/product-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	database.InitPostgres()

	r := gin.Default()
	routes.RegisterRoutes(r)
	r.Run(":8081")
}
