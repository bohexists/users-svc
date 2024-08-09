package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func createUserHandler(s Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		var u User
		if err := c.BindJSON(&u); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}
		s.CreateUser(u)
		c.JSON(http.StatusOK, u)
	}
}

func getUserHandler(s Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		u, err := s.GetUser(id)
		if err != nil || u == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusOK, u)
	}
}

func updateUserHandler(s Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var u User
		if err := c.BindJSON(&u); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}
		s.UpdateUser(id, u)
		c.JSON(http.StatusOK, u)
	}
}

func deleteUserHandler(s Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		s.DeleteUser(id)
		c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
	}
}

func getAllUsersHandler(s Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, _ := s.GetAllUsers()
		c.JSON(http.StatusOK, users)
	}
}
