package main

import (
	"user-service/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/health", handler.HealthCheck)

	r.Run(":8081")
}
