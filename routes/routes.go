package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jimtrung/amazon/handlers"
)

func SetupRoutes(app *fiber.App) {
	//Will make this look pretty later on
	//Products
	app.Get("api/products", handlers.GetProducts)
	app.Post("api/products/transfer", handlers.Transfer)
	app.Get("api/products/drop", handlers.DropTable)

	//Cart
	app.Get("api/cart", handlers.GetCart)
	app.Post("api/cart/add", handlers.AddToCart)
	app.Post("api/cart/update", handlers.UpdateCart)
	app.Post("api/cart/delete", handlers.DeleteFromCart)
}
