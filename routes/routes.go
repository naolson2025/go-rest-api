package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/naolson2025/go-rest-api/middlewares"
)

// we don't need to return anything
// bc we are modifying the server object directly
func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/signup", signup)
	server.POST("/login", login)

	// set up auth middleware on all these routes
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)
}