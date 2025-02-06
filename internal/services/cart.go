package services

import (
	"context"
	"strconv"

	"github.com/jimtrung/amazon/config"
	"github.com/jimtrung/amazon/internal/models"
)

func GetAllCarts() ([]models.CartWithItems, error) {
	query := `
		SELECT 
			ci.cart_item_id, ci.cart_id, ci.product_id, ci.quantity, ci.added_at,
			c.user_id, c.created_at, c.updated_at
		FROM cart_items ci
		JOIN carts c ON ci.cart_id = c.cart_id
	`

	rows, err := config.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cartItems []models.CartWithItems
	for rows.Next() {
		var cartItem models.CartWithItems
		err := rows.Scan(
			&cartItem.CartItemId,
			&cartItem.CartId,
			&cartItem.ProductId,
			&cartItem.Quantity,
			&cartItem.AddedAt,
			&cartItem.UserId,        // From carts table
			&cartItem.CartCreatedAt, // From carts table
			&cartItem.CartUpdatedAt, // From carts table
		)
		if err != nil {
			return nil, err
		}
		cartItems = append(cartItems, cartItem)
	}
	return cartItems, nil
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
