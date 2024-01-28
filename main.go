package main

import (
	"github.com/gin-gonic/gin"
	"github.com/naolson2025/go-rest-api/db"
	"github.com/naolson2025/go-rest-api/routes"
)

func main() {
	// db.InitDB()
	// server := gin.Default()

	// routes.RegisterRoutes(server)
	server := setupServer()

	server.Run(":8080")
}

func setupServer() *gin.Engine {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	return server
}
