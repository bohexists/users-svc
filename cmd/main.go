package main

import (
	"github.com/bohexists/users-svc/config"
	"github.com/bohexists/users-svc/controllers"
	"github.com/bohexists/users-svc/internal/middleware"
	"github.com/bohexists/users-svc/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig("config/config.yml")
	if err != nil {
		panic(err)
	}

	// Initialize Gin router
	r := gin.Default()

	// Apply all middleware
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.ErrorHandlingMiddleware())
	r.Use(middleware.RateLimiterMiddleware())

	// Initialize MongoDB repository
	repo := repository.NewMongoRepository(cfg.Mongo.URI, cfg.Mongo.Database, cfg.Mongo.Collection)

	// Initialize UserController with the repository (Controller layer)
	userController := controllers.NewUserController(repo)

	// Routes for HTML pages (MVC structure)
	r.GET("/users", userController.ShowUsers)             // Show list of users
	r.GET("/user/new", userController.NewUserForm)        // Show form to create new user
	r.POST("/user/create", userController.CreateUser)     // Handle user creation
	r.GET("/user/edit/:id", userController.EditUserForm)  // Show form to edit a user
	r.POST("/user/update/:id", userController.UpdateUser) // Handle user update
	r.GET("/user/delete/:id", userController.DeleteUser)  // Handle user deletion

	// Start the server
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
