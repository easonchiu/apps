package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type loginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func getAdminUsername() string {
	if v := os.Getenv("ADMIN_USERNAME"); v != "" {
		return v
	}
	return "admin"
}

func getAdminPassword() string {
	if v := os.Getenv("ADMIN_PASSWORD"); v != "" {
		return v
	}
	return "admin123"
}

func GetJWTSecret() []byte {
	if v := os.Getenv("JWT_SECRET"); v != "" {
		return []byte(v)
	}
	return []byte("ysgame-secret-key")
}

// AdminLogin handles admin login and returns a JWT token
func AdminLogin(c *gin.Context) {
	var req loginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Validate credentials against env vars
	if req.Username != getAdminUsername() || req.Password != getAdminPassword() {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Create JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": req.Username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
		"iat":      time.Now().Unix(),
	})

	tokenString, err := token.SignedString(GetJWTSecret())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
