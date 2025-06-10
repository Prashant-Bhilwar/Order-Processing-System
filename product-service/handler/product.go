package handler

import (
	"net/http"

	"github.com/prashant-bhilwar/order-processing-system/product-service/model"
	"github.com/prashant-bhilwar/order-processing-system/product-service/repository"

	"github.com/gin-gonic/gin"
)

// @Summary Create a new product
// @Description Add a new product to the catalog
// @Tags Products
// @Accept json
// @Produce json
// @Param product body model.Product true "Product to create"
// @Success 201 {object} model.Product
// @Failure 400 {object} map[string]string
// @failure 500 {object} map[string]string
// @Router /products [post]
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

// @Summary Get all products
// @Description Retrieve all products from the catalog
// @Tags Products
// @Produce json
// @Success 200 {array} model.Product
// @Failure 500 {object} map[string]string
// @Router /products [get]
func GetAllProducts(c *gin.Context) {
	products, err := repository.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}
	c.JSON(http.StatusOK, products)
}
