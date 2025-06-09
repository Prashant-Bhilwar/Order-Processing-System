package main

import (
	"user-service/config"
	"user-service/database"
	"user-service/redis"

	"user-service/handler"
	"user-service/middleware"

	"github.com/gin-gonic/gin"

	_ "user-service/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title User Service API
// @version 1.0
// @description This service handles user registration, login, and JWT auth
// @securityDefinitions.apiKey BearerAuth
// @in header
// @name Authorization
// @host localhost:8081
// @BasePath /
func main() {
	r := gin.Default()

	config.LoadConfig()
	database.InitPostgres()
	redis.InitRedis()

	authHandler := handler.NewAuthHandler(database.DB)

	r.GET("/health", handler.HealthCheck)
	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)
	r.POST("/refresh", authHandler.Refresh)

	authorized := r.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	authorized.GET("/me", handler.GetCurrentUser)
	authorized.POST("/logout", authHandler.Logout)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8081")
}
