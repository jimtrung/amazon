package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/jimtrung/amazon/config"
)

func main() {
	config.ConnectDB()
	app := fiber.New()

	app.Use(static.New("./amazon"))
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendFile("./amazon/amazon.html")
	})

	log.Fatal(app.Listen(":" + config.PORT))
}
