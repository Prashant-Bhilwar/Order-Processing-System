package main

import (
	"fmt"

	"github.com/prashant-bhilwar/order-processing-system/cart-service/config"
	"github.com/prashant-bhilwar/order-processing-system/cart-service/database"
)

func main() {
	config.LoadConfig()
	fmt.Println("Cart service started")
	database.InitPostgres()
}
