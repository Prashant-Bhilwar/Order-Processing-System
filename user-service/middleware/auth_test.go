package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/prashant-bhilwar/order-processing-system/user-service/token"

	"github.com/gin-gonic/gin"
)

func TestAuthMidddleware_ValidToken(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(AuthMiddleware())
	router.GET("/protected", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "ok"})
	})

	token, _ := token.GenerateAccessToken(999)
	req := httptest.NewRequest(http.MethodGet, "/protected", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}
