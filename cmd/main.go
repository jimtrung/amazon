package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jimtrung/amazon/api/routes"
	"github.com/jimtrung/amazon/config"
	"github.com/jimtrung/amazon/sql"
)

func main() {
	// Database setup
	config.ConnectDB()

	err := sql.SetupDatabase()
	if err != nil {
		log.Fatal("Failed to set up database:", err)
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.SetTrustedProxies([]string{"127.0.0.1"})

	// Routes handlers
	routes.SetupRoutes(r)

	// Auth setup

	// Middleware setup

	// Start the server
	r.Run(":" + config.PORT)
}
