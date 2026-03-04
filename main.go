package main

import (
	"example.com/rest-api/db"
	"example.com/rest-api/routes"
	"github.com/gin-gonic/gin"
)

// This is a simple Go REST API server buit w/ the Gin web framework

// Let's start by first handling an incoming request w/ gin
func main() {
	// Set up database
	db.InitDB()
	// Create Gin HTTP server w/ default middleware (logging + recovery)
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") // Run server on port 8080
}
