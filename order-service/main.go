package main

import (
	"fmt"

	"github.com/prashant-bhilwar/order-processing-system/order-service/config"
)

func main() {
	config.LoadConfig()
	fmt.Println("Order service started")
}
