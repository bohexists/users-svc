package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	storage := NewCacheStorage()

	r.POST("/user", createUserHandler(storage))
	r.GET("/user/:id", getUserHandler(storage))
	r.PUT("/user/:id", updateUserHandler(storage))
	r.DELETE("/user/:id", deleteUserHandler(storage))
	r.GET("/users", getAllUsersHandler(storage))

	r.Run(":8080")
}
