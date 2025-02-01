package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jimtrung/amazon/handlers"
	"github.com/jimtrung/amazon/internal/auth"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Products
	products := api.Group("/products")
	products.Get("/", handlers.GetProducts)       // ✅
	products.Post("/transfer", handlers.Transfer) // ✅
	products.Get("/drop", handlers.DropProducts)  // ✅

	// Cart
	cart := api.Group("/cart")
	cart.Get("/", handlers.GetCart)               // ✅
	cart.Post("/add", handlers.AddToCart)         // ✅
	cart.Post("/update", handlers.UpdateCart)     // ✅
	cart.Post("/delete", handlers.DeleteFromCart) // ✅
	cart.Get("/drop", handlers.DropCart)          // ✅

	//Authentication
	autho := app.Group("/auth")
	//Basic auth
	autho.Get("/basic", auth.BasicAuth)
}
