package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jimtrung/amazon/handlers"
	"github.com/jimtrung/amazon/internal/auth"
)

func SetupRoutes(app *gin.Engine) {
	api := app.Group("/api")

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

	//Authentication
	autho := app.Group("/auth")
	//Basic auth
	autho.GET("/basic", auth.BasicAuth)
}
