package routes

import (
	"example/web-service-gin/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitializeAuthRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	userHandler := handlers.NewUserHandler(db)

	r.GET("/users", userHandler.GetUsers)
	r.POST("/users", userHandler.CreateUser)
	r.PUT("/users/:id", userHandler.UpdateUser)
	r.DELETE("/users/:id", userHandler.DeleteUser)
	r.POST("/auth", userHandler.AuthenticateUser)

	return r
}
