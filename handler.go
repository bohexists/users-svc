package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// createUserHandler handles the creation of a new user
func createUserHandler(s Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		var u User
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

// getUserHandler handles retrieving a user
func getUserHandler(s Storage) gin.HandlerFunc {
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

// getAllUsersHandler handles retrieving all users
func getAllUsersHandler(s Storage) gin.HandlerFunc {
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

// updateUserHandler handles updating an existing user
func updateUserHandler(s Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var u User
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

// deleteUserHandler handles deleting a user
func deleteUserHandler(s Storage) gin.HandlerFunc {
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
