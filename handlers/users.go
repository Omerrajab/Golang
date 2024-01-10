// handlers/users.go

package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User is a struct representing a user
type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserHandler is a struct that handles user-related operations
type UserHandler struct {
	db *gorm.DB
}

// NewUserHandler creates a new UserHandler instance
func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{db: db}
}

// GetUsers handles GET request for fetching all users
func (h *UserHandler) GetUsers(c *gin.Context) {
	var users []User
	h.db.Find(&users)
	c.JSON(http.StatusOK, users)
}

// CreateUser handles POST request for creating a new user
func (h *UserHandler) CreateUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash the password before saving to the database
	enteredPassword := user.Password
	fmt.Println("enteredPassword", enteredPassword)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(enteredPassword), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)
	h.db.Create(&user)
	c.JSON(http.StatusCreated, user)
}

// UpdateUser handles PUT request for updating a user by ID
func (h *UserHandler) UpdateUser(c *gin.Context) {
	var user User
	id := c.Param("id")

	if err := h.db.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.db.Save(&user)
	c.JSON(http.StatusOK, user)
}

// DeleteUser handles DELETE request for deleting a user by ID
func (h *UserHandler) DeleteUser(c *gin.Context) {
	var user User
	id := c.Param("id")

	if err := h.db.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	h.db.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

