package database

import (
	"fmt"
	"log"

	"github.com/prashant-bhilwar/order-processing-system/cart-service/config"
	"github.com/prashant-bhilwar/order-processing-system/cart-service/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitPostgres() {
	var err error

	dsn := config.DBUrl
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Postgres connection error: %v", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed to get SQL DB: %v", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("Postgres ping failed: %v", err)
	}

	fmt.Println("PostgreSQL connected")

	err = DB.AutoMigrate(&model.CartItem{})
	if err != nil {
		log.Fatalf("failed to migrate CartItem: %v", err)
	}
	fmt.Println("CartItem table created")
}
