package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/jimtrung/amazon/api/routes"
	"github.com/jimtrung/amazon/config"
	"github.com/jimtrung/amazon/internal/middleware"
)

//	@title						Amazon
//	@version					1.0.0
//	@description				This project is a practice for Go Backend knowledge
//	@host						127.0.0.1:8080
//	@BasePath					/api
//	@schemes					http https
//	@securityDefinitions.basic	BasicAuth
//	@contact.name				jimtrung
//	@contact.email				nguyenhaitrung737@gmail.com

func main() {
	// Oauth
	middleware.NewAuth()

	// Database setup
	config.ConnectDB()

	// Config the server
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.SetTrustedProxies([]string{"127.0.0.1"})

	// Routes handlers
	routes.SetupRoutes(r)

	// Start the server
	fmt.Println("Server is running on port", config.PORT)
	r.Run("127.0.0.1:" + config.PORT)
}
