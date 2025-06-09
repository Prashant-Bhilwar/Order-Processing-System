package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUrl        string
	RedisAddress string
	JWTSecret    string
}

var AppConfig Config

func LoadConfig() {
	_ = godotenv.Load()

	AppConfig = Config{
		DBUrl:        getEnv("DB_URL", "postgres://user:password@localhost:5432/orderdb?sslmode=disable"),
		RedisAddress: getEnv("REDIS_ADDR", "localhost:6379"),
		JWTSecret:    getEnv("JWT_SECRET", "mysecret"),
	}

}

func getEnv(key string, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Printf("env: %s not set, using default", key)
		return fallback
	}
	return val
}
