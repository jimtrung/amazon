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
	products.GET("/", handlers.GetProducts)       // ✅
	products.POST("/transfer", handlers.Transfer) // ✅

	// Cart
	cart := api.Group("/cart")
	cart.GET("/", handlers.GetCart)                             // ✅
	cart.POST("/add", handlers.AddToCart)                       // ✅
	cart.PATCH("/update", handlers.UpdateCart)                  // ✅
	cart.DELETE("/delete/:product_id", handlers.DeleteFromCart) // ✅

	//User
	users := api.Group("/users")
	users.GET("/", handlers.GetUsers)
	users.POST("/signup", handlers.Signup)
	users.POST("/login", handlers.Login)

	// Admin
	protected := r.Group("/protected")
	protected.Use(middleware.BasicAuthMiddleware())
	{
		protected.POST("/auth", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Authourized"})
		})
		protected.DELETE("/delete/:user_id", handlers.DeleteUser)
		protected.DELETE("/drop-products", handlers.DropProducts) // ✅
		protected.DELETE("/drop-users", handlers.DropUsers)
		protected.DELETE("/drop-cart", handlers.DropCart) // ✅
	}

	// Authorization
	r.GET("/auth/:provider", middleware.BeginAuthProviderCallback)
	r.GET("/auth/:provider/callback", middleware.GetAuthCallBackFunction)

	//Serve static file
	r.StaticFile("/", "login.html")
}
