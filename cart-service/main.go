package main

import (
	"fmt"

	"github.com/prashant-bhilwar/order-processing-system/cart-service/config"
)

func main() {
	config.LoadConfig()
	fmt.Println("Cart service started")
}
