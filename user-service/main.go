package main

import (
	"user-service/database"
	"user-service/handler"

	"github.com/gin-gonic/gin"

	_ "user-service/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title User Service API
// @version 1.0
// @description This service handles user registration, login, and JWT auth
// @host localhost:8081
// @BasePath /
func main() {
	r := gin.Default()

	database.InitPostgres()

	authHandler := handler.NewAuthHandler(database.DB)

	r.GET("/health", handler.HealthCheck)
	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)
	r.POST("/refresh", authHandler.Refresh)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8081")
}
