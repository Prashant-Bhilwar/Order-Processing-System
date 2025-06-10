package main

import (
	"log"

	"github.com/prashant-bhilwar/order-processing-system/product-service/config"
	"github.com/prashant-bhilwar/order-processing-system/product-service/database"

	"github.com/prashant-bhilwar/order-processing-system/product-service/routes"

	"github.com/gin-gonic/gin"

	_ "github.com/prashant-bhilwar/order-processing-system/product-service/docs"

	"net"

	"github.com/prashant-bhilwar/order-processing-system/product-service/gen"
	productGrpc "github.com/prashant-bhilwar/order-processing-system/product-service/grpc"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/grpc"
)

// @title Product Service API
// @version 1.0
// @description REST APIs for managing products
// @host localhost:8081
// @BasePath /
func main() {
	config.LoadConfig()
	database.InitPostgres()

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routes.RegisterRoutes(r)

	go func() {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}

		grpcServer := grpc.NewServer()
		gen.RegisterProductServiceServer(grpcServer, &productGrpc.ProductGRPCServer{})

		log.Println("gRPC server started on :50051")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve ; %v", err)
		}
	}()

	r.Run(":8081")
}
