package handler

import (
	"net/http"
	"strconv"

	"github.com/prashant-bhilwar/order-processing-system/cart-service/model"
	"github.com/prashant-bhilwar/order-processing-system/cart-service/repository"

	"github.com/gin-gonic/gin"
)

// @Summary Add item to cart
// @Description Adds a product to user's cart
// @Tags Cart
// @Accept json
// @Produce json
// @Param item body model.CartItem true "Cart item"
// @Success 201 {object} model.CartItem
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /cart [post]
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

// @Summary Remove item from cart
// @Description Removes a specific product from user's cart
// @Tags Cart
// @Produce json
// @Param userId path int true "User ID"
// @Param productId path int true "Product ID"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /cart/{userId}/{productId} [delete]
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

// @Summary Get user's cart
// @Description Fetch all cart items for a given user
// @Tags Cart
// @Produce json
// @Param userId path int true "User ID"
// @Success 200 {array} model.CartItem
// @Failure 500 {object} map[string]string
// @Router /cart/{userid} [get]
func GetCart(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Param("userId"), 10, 64)

	items, err := repository.GetCartItemsByUserID(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch cart"})
		return
	}

	c.JSON(http.StatusOK, items)
}
