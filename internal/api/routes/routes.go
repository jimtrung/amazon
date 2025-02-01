package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jimtrung/amazon/api/middleware"
	"github.com/jimtrung/amazon/handlers"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")

	// Products
	products := api.Group("/products")
	products.GET("/", handlers.GetProducts)       // ✅
	products.POST("/transfer", handlers.Transfer) // ✅
	products.GET("/drop", handlers.DropProducts)  // ✅

	// Cart
	cart := api.Group("/cart")
	cart.GET("/", handlers.GetCart)               // ✅
	cart.POST("/add", handlers.AddToCart)         // ✅
	cart.POST("/update", handlers.UpdateCart)     // ✅
	cart.POST("/delete", handlers.DeleteFromCart) // ✅
	cart.GET("/drop", handlers.DropCart)          // ✅

	// Authentication
	protected := r.Group("/protected")
	protected.Use(middleware.BasicAuthMiddleware())
	{
		protected.GET("/auth", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Authourized"})
		})
	}

	//Basic auth
}
