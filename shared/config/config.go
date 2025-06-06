package config

import (
	"log"
	"os"
)

type Config struct {
	DBUrl     string
	JWTSecret string
}

func LoadConfig() *Config {
	cfg := &Config{
		DBUrl:     getEnv("DB_URL", "postgres://user:pass@localhost:5432/app"),
		JWTSecret: getEnv("JWT_SECRET", "defaultsecret"),
	}
	return cfg
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	log.Printf("Warning: using fallback value for %s", key)
	return fallback
}
