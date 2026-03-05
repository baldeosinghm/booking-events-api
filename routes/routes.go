package routes

import "github.com/gin-gonic/gin"

// Responsible for registering event routes

func RegisterRoutes(server *gin.Engine) {
	// User Gin HTTP server to register routes
	server.GET("/events", getEvents) // Create "/events" endpoint; make getEvents the endpoint handler
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvent)
	server.PUT("/events/:id", updateEvent)
	server.DELETE("/events/:id", deleteEvent)
	server.POST("/signup", signup)
	server.POST("/login", login)
}
