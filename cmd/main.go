package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jimtrung/amazon/internal/api/routes"
	"github.com/jimtrung/amazon/internal/config"
	sql "github.com/jimtrung/amazon/internal/database"
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

	// Start the server
	r.Run(":" + config.PORT)
}
