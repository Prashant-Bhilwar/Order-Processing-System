package handler

import (
	"net/http"
	"strconv"

	"github.com/prashant-bhilwar/order-processing-system/cart-service/model"
	"github.com/prashant-bhilwar/order-processing-system/cart-service/repository"

	"github.com/gin-gonic/gin"
)

func AddToCart(c *gin.Context) {
	var item model.CartItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := repository.AddToCart(&item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add to cart"})
		return
	}

	c.JSON(http.StatusCreated, item)
}

func RemoveFromCart(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Param("userId"), 10, 64)
	productID, _ := strconv.ParseUint(c.Param("productId"), 10, 64)

	err := repository.RemoveFromCart(uint(userID), uint(productID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item removed"})
}

func GetCart(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Param("userId"), 10, 64)

	items, err := repository.GetCartItemsByUserID(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch cart"})
		return
	}

	c.JSON(http.StatusOK, items)
}
