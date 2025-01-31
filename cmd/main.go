package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/jimtrung/amazon/config"
	"github.com/jimtrung/amazon/sql"
)

func main() {
	config.ConnectDB()
	err := sql.SetupDatabase()
	if err != nil {
		log.Fatal("Failed to set up database:", err)
	}
	app := fiber.New()

	app.Use(static.New("./amazon"))
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendFile("./amazon/amazon.html")
	})

	log.Fatal(app.Listen(":" + config.PORT))
}
