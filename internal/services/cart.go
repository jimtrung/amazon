package services

import (
	"context"
	"strconv"

	"github.com/jimtrung/amazon/config"
	"github.com/jimtrung/amazon/internal/models"
)

func GetAllCarts() ([]models.Cart, error) {
	query := `
		SELECT * FROM carts;
	`

	rows, err := config.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var carts []models.Cart
	for rows.Next() {
		var cart models.Cart
		err := rows.Scan(
			&cart.CartId,
			&cart.UserId,
			&cart.CreatedAt,
			&cart.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		carts = append(carts, cart)
	}
	return carts, nil
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
