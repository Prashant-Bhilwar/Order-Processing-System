package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/prashant-bhilwar/order-processing-system/order-service/gen"
	"github.com/prashant-bhilwar/order-processing-system/order-service/grpc"
	"github.com/prashant-bhilwar/order-processing-system/order-service/model"
	"github.com/prashant-bhilwar/order-processing-system/order-service/mq"
	"github.com/prashant-bhilwar/order-processing-system/order-service/repository"
	"github.com/rabbitmq/amqp091-go"

	"github.com/gin-gonic/gin"
)

// @Summary Create a new order
// @Description Place a new order for a user
// @Tags Orders
// @Accept json
// @Produce json
// @Param order body model.Order true "order payload"
// @Success 201 {object} model.Order
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /orders [post]
func CreateOrder(c *gin.Context) {
	var order model.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	order.Status = "pending"
	order.CreatedAt = time.Now()

	log.Println("Calling product-service with ProductID:", order.ProductID)
	product, err := grpc.ProductClient.GetProductById(context.Background(), &gen.ProductId{
		Id: int32(order.ProductID),
	})
	if err != nil {
		log.Println("gRPC error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get product"})
		return
	}

	order.TotalPrice = float64(product.Price)

	if err := repository.CreateOrder(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create order"})
		return
	}

	eventData, _ := json.Marshal(order)
	mq.Channel.PublishWithContext(c, "", "order_events", false, false, amqp091.Publishing{
		ContentType: "application/json",
		Body:        eventData,
	})

	c.JSON(http.StatusCreated, order)
}

// @Summary Get all orders by user
// @Description Retrieve all orders placed by a specific user
// @Tags Orders
// @Produce json
// @Param userId path int true "User ID"
// @Success 200 {array} model.Order
// @Failure 500 {object} map[string]string
// @Router /orders/{userId} [get]
func GetOrders(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Param("userId"), 10, 64)

	orders, err := repository.GetOrdersByUserID(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get orders"})
		return
	}

	c.JSON(http.StatusOK, orders)
}
