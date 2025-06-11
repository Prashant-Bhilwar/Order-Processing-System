package database

import (
	"fmt"
	"log"

	"github.com/prashant-bhilwar/order-processing-system/order-service/config"
	"github.com/prashant-bhilwar/order-processing-system/order-service/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitPostgres() {
	var err error
	DB, err := gorm.Open(postgres.Open(config.DBUrl), &gorm.Config{})
	if err != nil {
		log.Fatalf("Postgres connection error: %v", err)
	}

	err = DB.AutoMigrate(&model.Order{})
	if err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}

	fmt.Println("PostgreSQl connected and Order table created")
}
