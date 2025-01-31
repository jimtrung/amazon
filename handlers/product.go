package handlers

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/jimtrung/amazon/config"
	"github.com/jimtrung/amazon/models"
)

func GetProducts(c fiber.Ctx) error {
	rows, err := config.DB.Query(context.Background(), "SELECT * FROM products")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Image,
			&product.Rating.Stars,
			&product.Rating.Count,
			&product.PriceCents,
			&product.Keywords,
		)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		products = append(products, product)
	}

	return c.JSON(products)
}

func Transfer(c fiber.Ctx) error {
	var products []models.Product

	if err := c.Bind().JSON(&products); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	for _, product := range products {
		_, err := config.DB.Exec(
			context.Background(),
			"INSERT INTO products VALUES ($1, $2, $3, $4, $5, $6, $7)",
			product.Id, product.Name, product.Image,
			product.Rating.Stars, product.Rating.Count,
			product.PriceCents, product.Keywords,
		)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
	}

	return c.JSON(fiber.Map{"message": "Transfer successfully"})
}
