package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prashant-bhilwar/order-processing-system/cart-service/config"
	"github.com/prashant-bhilwar/order-processing-system/cart-service/database"
	"github.com/prashant-bhilwar/order-processing-system/cart-service/routes"
)

func main() {
	config.LoadConfig()
	database.InitPostgres()

	r := gin.Default()
	routes.SetupRoutes(r)

	r.Run(":8083")
}
