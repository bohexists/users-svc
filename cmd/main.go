package main

import (
	"github.com/bohexists/users-svc/internal"
	"github.com/gin-gonic/gin"
)

func main() {

	// Initialize Gin router
	r := gin.Default()

	// Apply all middleware
	r.Use(internal.CORSMiddleware())
	r.Use(internal.ErrorHandlingMiddleware())
	r.Use(internal.RateLimiterMiddleware())

	// Initialize CacheStorage
	storage := internal.NewCacheStorage()

	// Routes.
	r.POST("/user", internal.CreateUserHandler(storage))
	r.GET("/user/:id", internal.GetUserHandler(storage))
	r.PUT("/user/:id", internal.UpdateUserHandler(storage))
	r.DELETE("/user/:id", internal.DeleteUserHandler(storage))
	r.GET("/users", internal.GetAllUsersHandler(storage))
	r.GET("/user/search", internal.SearchUserByEmailHandler(storage))
	// Start the server
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
