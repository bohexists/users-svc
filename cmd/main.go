package main

import (
	"github.com/bohexists/users-svc/config"
	"github.com/bohexists/users-svc/controllers"
	"github.com/bohexists/users-svc/internal/middleware"
	"github.com/bohexists/users-svc/repository"
	"github.com/gin-gonic/gin"
	"log"
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

	// Log the database configuration
	log.Printf("Connecting to MongoDB database: %s", cfg.Mongo.Database)

	// Initialize MongoDB repository
	repo := repository.NewMongoRepository(cfg.Mongo.URI, cfg.Mongo.Database, cfg.Mongo.Collection)

	// Initialize UserController with the repository (Controller layer)
	userController := controllers.NewUserController(repo)

	// Routes for HTML pages (MVC structure)
	r.GET("/users", userController.ShowUsers)                    // Show list of users
	r.GET("/repositorys/new", userController.NewUserForm)        // Show form to create new repositorys
	r.POST("/repositorys/create", userController.CreateUser)     // Handle repositorys creation
	r.GET("/repositorys/edit/:id", userController.EditUserForm)  // Show form to edit a repositorys
	r.POST("/repositorys/update/:id", userController.UpdateUser) // Handle repositorys update
	r.GET("/repositorys/delete/:id", userController.DeleteUser)  // Handle repositorys deletion

	// Start the server
	if err := r.Run(":8080"); err != nil {
		return
	}
}
