package services

import (
	"context"
	"strconv"

	"github.com/jimtrung/amazon/config"
	"github.com/jimtrung/amazon/internal/models"
)

func GetAllCarts() ([]models.CartItem, error) {
	rows, err := config.DB.Query(context.Background(), "SELECT * FROM carts")
	if err != nil {
		return []models.CartItem{}, err
	}
	defer rows.Close()

	var cart []models.CartItem
	for rows.Next() {
		var cartItem models.CartItem
		err := rows.Scan(
			&cartItem.CartItemId,
			&cartItem.CartId,
			&cartItem.ProductId,
			&cartItem.Quantity,
			&cartItem.AddedAt,
		)
		if err != nil {
			return []models.CartItem{}, err
		}
		cart = append(cart, cartItem)
	}
	return cart, nil
}

func AddToCart(cartId int, productId string, quantity int) error {
	_, err := config.DB.Exec(
		context.Background(),
		"SELECT add_to_cart($1, $2, $3);",
		cartId, productId, quantity,
	)

	return err
}

func UpdateCartItemQuantity(cartId int, productId string, quantity int) error {
	_, err := config.DB.Exec(
		context.Background(),
		"SELECT update_cart($1, $2, $3);",
		cartId, productId, quantity,
	)

	return err
}

func DeleteFromCart(cartId int, productId string) error {
	_, err := config.DB.Exec(
		context.Background(),
		"SELECT delete_from_cart($1, $2)",
		cartId, productId,
	)

	return err
}

func StringToInt(str string) (int, error) {
	return strconv.Atoi(str)
}
