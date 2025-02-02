package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jimtrung/amazon/internal/api/handlers"
	"github.com/jimtrung/amazon/internal/api/middleware"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")

	// Products
	products := api.Group("/products")
	products.GET("/", handlers.GetProducts)         // ✅
	products.POST("/transfer", handlers.Transfer)   // ✅
	products.DELETE("/drop", handlers.DropProducts) // ✅

	// Cart
	cart := api.Group("/cart")
	cart.GET("/", handlers.GetCart)                             // ✅
	cart.POST("/add", handlers.AddToCart)                       // ✅
	cart.PATCH("/update", handlers.UpdateCart)                  // ✅
	cart.DELETE("/delete/:product_id", handlers.DeleteFromCart) // ✅
	cart.DELETE("/drop", handlers.DropCart)                     // ✅

	//User
	users := api.Group("/users")
	users.GET("/", handlers.GetUsers)
	users.POST("/add", handlers.AddUser)
	users.DELETE("/delete/:user_id", handlers.DeleteUser)
	users.DELETE("/drop", handlers.DropUsers)

	// Authorization
	protected := r.Group("/protected")
	protected.Use(middleware.BasicAuthMiddleware())
	{
		protected.POST("/auth", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Authourized"})
		})
	}
	r.GET("/auth/:provider", middleware.BeginAuthProviderCallback)
	r.GET("/auth/:provider/callback", middleware.GetAuthCallBackFunction)

	//Serve static file
	r.StaticFile("/", "login.html")
}
