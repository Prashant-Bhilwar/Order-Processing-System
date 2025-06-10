package handler

import (
	"net/http"

	"github.com/prashant-bhilwar/order-processing-system/product-service/model"
	"github.com/prashant-bhilwar/order-processing-system/product-service/repository"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var p model.Product
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := repository.CreateProduct(&p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusCreated, p)
}

func GetAllProducts(c *gin.Context) {
	products, err := repository.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}
	c.JSON(http.StatusOK, products)
}
