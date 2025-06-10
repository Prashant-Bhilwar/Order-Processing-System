package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUrl string
}

var AppConfig Config

func LoadConfig() {
	_ = godotenv.Load()
	AppConfig = Config{
		DBUrl: getEnv("DB_URL", "postgres://user:password@localhost:5432/orderdb?sslmode=disable"),
	}
}

func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Printf("env: %s not set, using default", key)
		return fallback
	}
	return val
}
