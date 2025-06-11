package grpc

import (
	"log"

	"github.com/prashant-bhilwar/order-processing-system/order-service/gen"

	"google.golang.org/grpc"
)

var ProductClient gen.ProductServiceClient

func InitProductGRPC() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to product-service gRPC: %v", err)
	}
	ProductClient = gen.NewProductServiceClient(conn)
}
