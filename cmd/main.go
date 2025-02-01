package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jimtrung/amazon/config"
	"github.com/jimtrung/amazon/routes"
	"github.com/jimtrung/amazon/sql"
)

func main() {
	// Database setup
	config.ConnectDB()

	err := sql.SetupDatabase()
	if err != nil {
		log.Fatal("Failed to set up database:", err)
	}

	// Create app
	app := gin.New()
	gin.SetMode(gin.ReleaseMode)
	app.SetTrustedProxies([]string{"127.0.0.1"})

	// Routes handlers
	routes.SetupRoutes(app)

	// Auth setup

	// Middleware setup

	// Start the server
	app.Run(":" + config.PORT)
}
