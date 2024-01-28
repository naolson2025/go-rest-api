package routes

import "github.com/gin-gonic/gin"

// we don't need to return anything
// bc we are modifying the server object directly
func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvent)
	server.PUT("/events/:id", updateEvent)
}