package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/prashant-bhilwar/order-processing-system/product-service/config"
)

var DB *sql.DB

func InitPostgres() {
	var err error
	DB, err := sql.Open("postgres", config.AppConfig.DBUrl)
	if err != nil {
		log.Fatalf("Postgres connection error: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Postgres ping failed: %v", err)
	}

	fmt.Printf("Postgres connected")
}
