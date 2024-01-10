package main

import (
	"example/web-service-gin/handlers"
	"example/web-service-gin/helpers/dbHelpers"
	middleware "example/web-service-gin/middlewares"
	"example/web-service-gin/routes"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	// Connect to the PostgreSQL database
	db, err := dbHelpers.InitDB()
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return
	}

	// Auto Migrate the User model
	db.AutoMigrate(&handlers.User{})

	r := gin.Default()
	r.Use(middleware.AuthMiddleware)
	routes.InitializeUserRoutes(db)
	r.GET("/public", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "This is a public endpoint"})
	})

	r.Use(middleware.AuthorizationMiddleware)

	r.GET("/private", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "This is a private endpoint"})
	})

	routes.InitializeUserRoutes(db)
	r.Run(":8085")

}
