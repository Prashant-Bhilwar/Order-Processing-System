package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/prashant-bhilwar/order-processing-system/user-service/middleware"
	"github.com/prashant-bhilwar/order-processing-system/user-service/token"

	"github.com/gin-gonic/gin"
)

func TestGetcurrentuser(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.Use(middleware.AuthMiddleware())

	router.GET("/me", GetCurrentUser)

	// generate token for test user ID (e.g. 42)
	accesstoken, _ := token.GenerateAccessToken(42)

	req := httptest.NewRequest(http.MethodGet, "/me", nil)
	req.Header.Set("Authorization", "Bearer "+accesstoken)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected 200 OK, got %d", w.Code)
	}
}
