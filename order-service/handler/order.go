package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/prashant-bhilwar/order-processing-system/order-service/model"
	"github.com/prashant-bhilwar/order-processing-system/order-service/repository"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	var order model.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	order.Status = "pending"
	order.CreatedAt = time.Now()

	if err := repository.CreateOrder(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create order"})
		return
	}

	c.JSON(http.StatusCreated, order)
}

func GetOrders(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Param("userId"), 10, 64)

	orders, err := repository.GetOrdersByUserID(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get orders"})
		return
	}

	c.JSON(http.StatusOK, orders)
}
