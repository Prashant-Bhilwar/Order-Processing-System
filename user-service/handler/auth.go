package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"user-service/model"
	rdb "user-service/redis"
	"user-service/repository"
	"user-service/token"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	UserRepo *repository.UserRepository
}

func NewAuthHandler(db *sql.DB) *AuthHandler {
	return &AuthHandler{UserRepo: repository.NewUserRepo(db)}
}

// @Summary Register a new user
// @Tags Auth
// @Accept json
// @Produce json
// @param user body model.User true "User registration input"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var input model.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	input.Password = string(hashedPassword)

	if err := h.UserRepo.CreateUser(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": input.ID, "email": input.Email})
}

// @Summary Login user and return JWT tokens
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body model.User true "User login input"
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var input model.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	user, err := h.UserRepo.GetByEmail(input.Email)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	accessToken, _ := token.GenerateAccessToken(user.ID)
	refreshToken, _ := token.GenerateRefreshToken(user.ID)

	key := fmt.Sprintf("refresh:%d", user.ID)
	rdb.Rdb.Set(rdb.Ctx, key, refreshToken, 7*24*time.Hour)

	c.JSON(http.StatusOK, gin.H{
		"message":       "login successful",
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

// @Summary Refresh access token
// @Tags Auth
// @Accept json
// @Produce json
// @Param token body map[string]string true "Refresh token input"
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /refresh [post]
func (h *AuthHandler) Refresh(c *gin.Context) {
	var body struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := c.ShouldBindJSON(&body); err != nil || body.RefreshToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	claims, err := token.ParseToken(body.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired refresh token"})
		return
	}

	key := fmt.Sprintf("refresh:%d", claims.UserID)
	storedToken, err := rdb.Rdb.Get(rdb.Ctx, key).Result()
	if err != nil || storedToken != body.RefreshToken {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "refresh token not recognized"})
		return
	}

	accessToken, _ := token.GenerateRefreshToken(claims.UserID)

	c.JSON(http.StatusOK, gin.H{
		"access_token": accessToken,
	})
}

// @Summary Logout user
// @Description Revoke refresh token from Redis and end the session
// @Tags Auth
// @Security BearerAuth
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /logout [post]
func (h *AuthHandler) Logout(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	key := fmt.Sprintf("refresh:%d", userID)
	err := rdb.Rdb.Del(rdb.Ctx, key).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to revoke token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "logout successful"})
}
