package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	// Initialize Gin router
	r := gin.Default()

	// Configure CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Initialize CacheStorage
	storage := NewCacheStorage()
	// Routes.
	r.POST("/user", createUserHandler(storage))
	r.GET("/user/:id", getUserHandler(storage))
	r.PUT("/user/:id", updateUserHandler(storage))
	r.DELETE("/user/:id", deleteUserHandler(storage))
	r.GET("/users", getAllUsersHandler(storage))
	// Start the server
	r.Run(":8080")
}
