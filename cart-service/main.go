package main

import (
	"log"
	"net"

	"github.com/gin-gonic/gin"
	"github.com/prashant-bhilwar/order-processing-system/cart-service/config"
	"github.com/prashant-bhilwar/order-processing-system/cart-service/database"
	"github.com/prashant-bhilwar/order-processing-system/cart-service/gen"
	cartGrpc "github.com/prashant-bhilwar/order-processing-system/cart-service/grpc"
	"github.com/prashant-bhilwar/order-processing-system/cart-service/routes"
	"google.golang.org/grpc"

	_ "github.com/prashant-bhilwar/order-processing-system/cart-service/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Cart Service API
// @version 1.0
// @description API for managing user cart operations
// @host localhost:8083
// @BasePath /
// @schemes http
func main() {
	config.LoadConfig()
	database.InitPostgres()

	r := gin.Default()
	routes.SetupRoutes(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	go func() {
		lis, err := net.Listen("tcp", ":50052")
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}

		grpcServer := grpc.NewServer()
		gen.RegisterCartServiceServer(grpcServer, &cartGrpc.CartGRPCServer{})

		log.Println("gRPC server started on :50052")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	r.Run(":8083")
}
