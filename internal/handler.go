package internal

import (
	"net/http"

	"github.com/bohexists/users-svc/models"
	"github.com/gin-gonic/gin"
)

// CreateUserHandler handles the creation of a new user
func CreateUserHandler(s Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		var u models.User
		if err := c.BindJSON(&u); err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{
				Code:    http.StatusBadRequest,
				Message: "Invalid request data",
				Details: err.Error(),
			})
			return
		}

		id, err := s.CreateUser(u)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: "Failed to create user",
				Details: err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{"id": id})
	}
}

// GetUserHandler handles retrieving a user
func GetUserHandler(s Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		u, err := s.GetUser(id)
		if err != nil || u == nil {
			c.JSON(http.StatusNotFound, ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: "Failed to retrieve user",
				Details: err.Error(),
			})
			return
		}

		if u == nil {
			c.JSON(http.StatusNotFound, ErrorResponse{
				Code:    http.StatusNotFound,
				Message: "User not found",
			})
			return
		}

		c.JSON(http.StatusOK, u)
	}
}

// GetAllUsersHandler handles retrieving all users
func GetAllUsersHandler(s Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := s.GetAllUsers()
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: "Failed to retrieve users",
				Details: err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, users)
	}
}

// UpdateUserHandler handles updating an existing user
func UpdateUserHandler(s Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var u models.User
		if err := c.BindJSON(&u); err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{
				Code:    http.StatusBadRequest,
				Message: "Invalid request data",
				Details: err.Error(),
			})
			return
		}

		if err := s.UpdateUser(id, u); err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: "Failed to update user",
				Details: err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, u)
	}
}

// DeleteUserHandler handles deleting a user
func DeleteUserHandler(s Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if err := s.DeleteUser(id); err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: "Failed to delete user",
				Details: err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
	}
}

// SearchUserByEmailHandler handles searching a user by email
func SearchUserByEmailHandler(s Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.Query("email")
		if email == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email query parameter is required"})
			return
		}

		users, err := s.GetAllUsers()
		if err != nil {
			c.Error(err).SetType(gin.ErrorTypePrivate)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search user by email"})
			return
		}

		for _, user := range users {
			if user.Email == email {
				c.JSON(http.StatusOK, user)
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	}
}
