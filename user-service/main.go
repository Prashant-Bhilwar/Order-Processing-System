package main

import (
	"user-service/database"
	"user-service/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	database.InitPostgres()

	authHandler := handler.NewAuthHandler(database.DB)

	r.GET("/health", handler.HealthCheck)
	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)
	r.POST("/refresh", authHandler.Refresh)

	r.Run(":8081")
}
