package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/jimtrung/amazon/api/routes"
	"github.com/jimtrung/amazon/config"
	"github.com/jimtrung/amazon/internal/database"
	"github.com/jimtrung/amazon/internal/middleware"
)

func main() {
	// Oauth
	middleware.NewAuth()

	// Database setup
	config.ConnectDB()
	err := database.SetupDatabase()
	if err != nil {
		log.Fatal("Failed to set up database:", err)
	}

	// Config the server
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.SetTrustedProxies([]string{"127.0.0.1"})

	// Routes handlers
	routes.SetupRoutes(r)

	// Start the server
	r.Run(":" + config.PORT)
}
