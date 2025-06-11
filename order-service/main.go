package main

import (
	"fmt"

	"github.com/prashant-bhilwar/order-processing-system/order-service/config"
	"github.com/prashant-bhilwar/order-processing-system/order-service/database"
)

func main() {
	config.LoadConfig()
	database.InitPostgres()
	fmt.Println("Order service started")
}
