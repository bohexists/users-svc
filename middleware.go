package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// ErrorResponse structure of the response when an error
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

// ErrorHandlingMiddleware centralized error handling
func ErrorHandlingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Check if there are any errors that occurred in the handler
		if len(c.Errors) > 0 {
			// Get the last error
			err := c.Errors.Last()

			// Return a response with detailed information about the error
			c.JSON(-1, ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: "Internal Server Error",
				Details: err.Error(),
			})
			c.Abort()
		}
	}
}

// CORSMiddleware configures CORS settings
func CORSMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}
