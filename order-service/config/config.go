package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var DBUrl string

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	DBUrl = os.Getenv("DB_URL")
	if DBUrl == "" {
		log.Fatalf("DB_URL is not set")
	}
}
