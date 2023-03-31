/*
@Time : 2023/3/31 12:00
@Author : sc-52766
@File : http_server.go
@Software: GoLand
*/
package api

// Import the necessary packages
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Define a struct to hold the necessary information for request validation
type Request struct {
	// Define the necessary fields for request validation
}

// Define a function to handle panics and recover gracefully
func handlePanic() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Handle the panic and return an error response
			}
		}()
		c.Next()
	}
}

// Define a function to register routes and handle requests
func RegisterRoutes() *gin.Engine {
	// Create a new gin engine
	r := gin.New()

	// Use the handlePanic middleware to handle panics and recover gracefully
	r.Use(handlePanic())

	// Register routes and handle requests
	r.GET("/", func(c *gin.Context) {
		// Handle the request and return a response
	})

	// Return the gin engine
	return r
}

// Call the RegisterRoutes function to start the server
func StartServer() {
	// Call the RegisterRoutes function to register routes and handle requests
	r := RegisterRoutes()

	// Start the server
	if err := r.Run(":8080"); err != nil && err != http.ErrServerClosed {
		// Handle the error and return an error response
	}
}
