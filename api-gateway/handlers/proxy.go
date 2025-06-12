package handlers

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
)

func createReverseProxy(target string, c *gin.Context) {
	url, err := url.Parse(target)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "invalid target"})
		return
	}
	proxy := httputil.NewSingleHostReverseProxy(url)
	c.Request.URL.Path = c.Param("proxyPath")
	proxy.ServeHTTP(c.Writer, c.Request)
}

func UserProxy(c *gin.Context) {
	createReverseProxy(os.Getenv("USER_SERVICE_URL"), c)
}

func ProductProxy(c *gin.Context) {
	createReverseProxy(os.Getenv("PRODUCT_SERVICE_URL"), c)
}

func OrderProxy(c *gin.Context) {
	createReverseProxy(os.Getenv("ORDER_SERVICE_URL"), c)
}

func CartProxy(c *gin.Context) {
	createReverseProxy(os.Getenv("CART_SERVICE_URL"), c)
}
