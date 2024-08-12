package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Gin router
	r := gin.Default()

	// Apply all middleware
	r.Use(ErrorHandlingMiddleware())
	r.Use(CORSMiddleware())

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
