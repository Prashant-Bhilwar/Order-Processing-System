package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/prashant-bhilwar/order-processing-system/user-service/database"
	"github.com/prashant-bhilwar/order-processing-system/user-service/model"
	"github.com/prashant-bhilwar/order-processing-system/user-service/redis"

	"testing"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	database.InitPostgres()
	redis.InitRedis()
	handler := NewAuthHandler(database.DB)

	r := gin.Default()
	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)
	return r
}

func TestRegisterLogin(t *testing.T) {
	r := setupRouter()

	user := model.User{
		Username: "testuser",
		Email:    fmt.Sprintf("test+%d@example.com", time.Now().UnixNano()),
		Password: "testpass123",
	}
	body, _ := json.Marshal(user)

	// Test /register
	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	if resp.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", resp.Code)
	}

	// Test /login
	req2 := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
	req2.Header.Set("Content-Type", "application/json")
	resp2 := httptest.NewRecorder()
	r.ServeHTTP(resp2, req2)

	if resp2.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", resp2.Code)
	}
}
