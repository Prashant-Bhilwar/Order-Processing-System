package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/prashant-bhilwar/order-processing-system/api-gateway/routes"
	"github.com/prashant-bhilwar/order-processing-system/api-gateway/utils"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	utils.InitLogger()
	router := routes.SetupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Println("API Gateway running on port:", port)
	router.Run(":" + port)
}
