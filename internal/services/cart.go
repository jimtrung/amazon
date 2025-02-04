package services

import (
	"context"

	"github.com/jimtrung/amazon/config"
	"github.com/jimtrung/amazon/internal/models"
)

func GetAllCart() ([]models.CartItem, error) {
	rows, err := config.DB.Query(context.Background(), "SELECT * FROM cart")
	if err != nil {
		return []models.CartItem{}, err
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
			return []models.CartItem{}, err
		}
		cart = append(cart, cartItem)
	}
	return cart, nil
}

func AddToCart(productId string, quantity int) error {
	_, err := config.DB.Exec(
		context.Background(),
		"SELECT add_to_cart($1, $2);",
		productId, quantity,
	)
	if err != nil {
		return err
	}
	return nil
}

func UpdateCartItemQuantity(productId string, quantity int) error {
    _, err := config.DB.Exec(
        context.Background(),
        "SELECT update_cart($1, $2);",
        productId, quantity,
    )
    if err != nil {
        return err
    }
    return nil
}

func DeleteFromCart(productId string) error {
	_, err := config.DB.Exec(
		context.Background(),
		"SELECT delete_from_cart($1)",
		productId,
	)
	if err != nil {
        return err
	}
    return nil
}

func DropCart() error {
	dropTable := `
		DROP TABLE cart;
	`

	_, err := config.DB.Exec(
		context.Background(),
		dropTable,
	)
	if err != nil {
        return err
	}
    return nil
}
