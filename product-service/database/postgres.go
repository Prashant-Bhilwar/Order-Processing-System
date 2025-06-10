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
	DB, err = sql.Open("postgres", config.AppConfig.DBUrl)
	if err != nil {
		log.Fatalf("Postgres connection error: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Postgres ping failed: %v", err)
	}

	fmt.Println("Postgres connected")

	createTable := `
	CREATE TABLE IF NOT EXISTS products (
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL,
	description TEXT,
	price REAL NOT NULL,
	in_stock BOOLEAN NOT NULL DEFAULT true
	);`

	_, err = DB.Exec(createTable)
	if err != nil {
		log.Fatalf("Failed to create products table: %v", err)
	}

	fmt.Println("Products table checked/created")
}
