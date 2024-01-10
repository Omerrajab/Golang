package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)



func newAuthHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{db: db}
}


func (h *UserHandler) AuthenticateUser(c *gin.Context) {
	var loginRequest struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.authenticateUser(loginRequest.Email, loginRequest.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// You can generate a JWT token or set a session cookie here for authenticated users
	c.JSON(http.StatusOK, gin.H{"message": "Authentication successful", "user": user})
}

// authenticateUser checks if the provided password matches the stored hashed password
func (h *UserHandler) authenticateUser(email, password string) (*User, error) {
	var user User
	result := h.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	fmt.Println("result:", user, password)
	// Compare the stored hashed password with the provided password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		fmt.Println("error:", err)
		return nil, err // Passwords do not match
	}

	return &user, nil
}