package middleware

import (
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)
const secretKeyEnvVar = "SECRET_KEY"

func AuthMiddleware(c *gin.Context) {
	secretKey := os.Getenv(secretKeyEnvVar)
	if secretKey == "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Secret key not set"})
		return
	}

	tokenString, err := c.Cookie("token")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	claims := &jwt.StandardClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	// You can access claims.ID, claims.Subject, claims.ExpiresAt, etc.

	c.Next()
}
