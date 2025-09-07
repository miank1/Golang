package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func loggingMiddleware(c *gin.Context) {
	fmt.Printf("Request: %s %s\n", c.Request.Method, c.Request.URL.Path)
	c.String(200, "From Middleware!")
	fmt.Println("Request processed")
	// Call the next handler
	c.Next()

}

func main() {
	r := gin.Default()

	// Apply middleware globally
	r.Use(loggingMiddleware)

	// Define a route
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome to the Home Page!")
	})

	// Start the server
	r.Run(":8082")
}
