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
	app := gin.Default()

	// Routes handlers
	routes.SetupRoutes(app)

	// Auth setup

	// Middleware setup
	// app.Use(static.New("./amazon"))
	// app.Get("/", func(c fiber.Ctx) error {
	// 	return c.SendFile("./amazon/auth.html")
	// })

	// Start the server
	app.Run(":" + config.PORT)
}
