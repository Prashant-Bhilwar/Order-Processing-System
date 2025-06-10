package main

import (
	"github.com/prashant-bhilwar/order-processing-system/product-service/config"
	"github.com/prashant-bhilwar/order-processing-system/product-service/database"
)

func main() {
	config.LoadConfig()
	database.InitPostgres()
}
