package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// This is a simple Go REST API server buit w/ the Gin web framework

// Let's start by first handling an incoming request w/ gin
func main() {
	// Create Gin HTTP server w/ default middleware (logging + recovery)
	server := gin.Default()

	server.GET("/events", getEvents) // Create /events endpoint

	server.Run(":8080") // Run server on port 8080
}

func getEvents(context *gin.Context) {
	// Return JSON response
	context.JSON(http.StatusOK, gin.H{"message": "Hello!"})
}
