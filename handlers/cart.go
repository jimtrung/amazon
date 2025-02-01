package handlers

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/jimtrung/amazon/config"
	"github.com/jimtrung/amazon/models"
)

func GetCart(c fiber.Ctx) error {
	rows, err := config.DB.Query(context.Background(), "SELECT * FROM cart")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	defer rows.Close()

	var cart []models.CartItem
	for rows.Next() {
		var cartItem models.CartItem
		err := rows.Scan(
			&cartItem.ProductId,
			&cartItem.Quantity,
		)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		cart = append(cart, cartItem)
	}

	return c.JSON(cart)
}

// Add item to cart
func AddToCart(c fiber.Ctx) error {
	var cartItem models.CartItem
	if err := c.Bind().JSON(&cartItem); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	_, err := config.DB.Exec(
		context.Background(),
		"SELECT add_to_cart($1, $2);",
		cartItem.ProductId, cartItem.Quantity,
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Item added to cart"})
}

func UpdateCart(c fiber.Ctx) error {
	var cartItem models.CartItem
	if err := c.Bind().JSON(&cartItem); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	_, err := config.DB.Exec(
		context.Background(),
		"SELECT update_cart($1, $2)",
		cartItem.ProductId, cartItem.Quantity,
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Updated cart"})
}

func DeleteFromCart(c fiber.Ctx) error {
	var cartItem models.CartItem
	if err := c.Bind().JSON(&cartItem); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	_, err := config.DB.Exec(
		context.Background(),
		"DELETE FROM cart where product_id = $1",
		cartItem.ProductId,
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Item deleted to cart"})
}

func DropCart(c fiber.Ctx) error {
	dropTable := `
		DROP TABLE cart; 
	`

	_, err := config.DB.Exec(
		context.Background(),
		dropTable,
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Drop successfully"})
}
