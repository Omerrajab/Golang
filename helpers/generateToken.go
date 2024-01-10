package helpers

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const secretKeyEnvVar = "SECRET_KEY"

func generateToken() (string, error) {
	secretKey := os.Getenv(secretKeyEnvVar)
	if secretKey == "" {
		return "", fmt.Errorf("Secret key not set")
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() // Token expiration time (1 hour)
	// Add additional claims like user ID, username, roles, etc.
	// claims["sub"] = "user123"
	// claims["roles"] = []string{"admin", "user"}

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
