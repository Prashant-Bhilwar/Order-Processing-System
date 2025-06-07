package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitPostgres() {
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		dbURL = "postgres://user:password@localhost:5432/orderdb?sslmode=disable"
	}

	var err error
	DB, err = sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Postgres connection error: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Postgres ping failed: %v", err)
	}

	fmt.Println("Postgres connected")
}
