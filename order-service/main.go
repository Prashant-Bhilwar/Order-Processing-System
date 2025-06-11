package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prashant-bhilwar/order-processing-system/order-service/config"
	"github.com/prashant-bhilwar/order-processing-system/order-service/database"
	"github.com/prashant-bhilwar/order-processing-system/order-service/routes"
)

func main() {
	config.LoadConfig()
	database.InitPostgres()

	r := gin.Default()
	routes.SetupRoutes(r)

	r.Run(":8084")
}
