package controllers

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/bohexists/users-svc/models"
	"github.com/bohexists/users-svc/repository"
)

type UserController struct {
	repo repository.Repository
}

// NewUserController creates a new user controller
func NewUserController(repo repository.Repository) *UserController {
	return &UserController{repo: repo}
}

// ShowUsers renders the page with a list of all users
func (uc *UserController) ShowUsers(c *gin.Context) {
	users, err := uc.repo.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load users"})
		return
	}

	// Load HTML template
	tmpl, err := template.ParseFiles("views/users.html")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load template"})
		return
	}

	// Render the template with the user data
	tmpl.Execute(c.Writer, gin.H{"Users": users})
}

// NewUserForm renders the form for creating a new user
func (uc *UserController) NewUserForm(c *gin.Context) {
	// Load the HTML template for the form
	tmpl, err := template.ParseFiles("views/new_user.html")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load template"})
		return
	}

	// Render the form template
	tmpl.Execute(c.Writer, nil)
}

// CreateUser handles the form submission to create a new user
func (uc *UserController) CreateUser(c *gin.Context) {
	var u models.User

	// Bind the form data to the user model
	if err := c.Bind(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form data"})
		return
	}

	// Create a new user through the repository
	_, err := uc.repo.CreateUser(u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Redirect to the user list page after successful creation
	c.Redirect(http.StatusFound, "/users")
}

// EditUserForm renders the form for editing an existing user
func (uc *UserController) EditUserForm(c *gin.Context) {
	id := c.Param("id")

	// Retrieve user data from the repository
	user, err := uc.repo.GetUser(id)
	if err != nil || user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Load the HTML template for the edit form
	tmpl, err := template.ParseFiles("views/edit_user.html")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load template"})
		return
	}

	// Render the form with the current user data
	tmpl.Execute(c.Writer, gin.H{"User": user})
}

// UpdateUser handles form submission for updating an existing user
func (uc *UserController) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var u models.User

	// Bind the form data to the user model
	if err := c.Bind(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form data"})
		return
	}

	// Update the user in the repository
	err := uc.repo.UpdateUser(id, u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	// Redirect to the user list page after successful update
	c.Redirect(http.StatusFound, "/users")
}

// DeleteUser handles the deletion of a user
func (uc *UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	// Delete the user through the repository
	err := uc.repo.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	// Redirect to the user list page after successful deletion
	c.Redirect(http.StatusFound, "/users")
}
