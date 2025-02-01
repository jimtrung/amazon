package migrations

import (
	"context"
	"fmt"

	"github.com/jimtrung/amazon/config"
)

func CreateAddToCart() error {
	addToCart := `
		CREATE OR REPLACE FUNCTION add_to_cart(add_id UUID, add_quantity INT) 
		RETURNS VOID AS $$
		BEGIN
			INSERT INTO cart (product_id, quantity) 
			VALUES (add_id, add_quantity)
			ON CONFLICT (product_id) 
			DO UPDATE SET quantity = cart.quantity + add_quantity;
		END;
		$$ LANGUAGE plpgsql;
	`

	_, err := config.DB.Exec(
		context.Background(),
		addToCart,
	)
	if err != nil {
		return err
	}

	fmt.Println("003_add_to_cart(1/1) - Add to cart function create successfully")
	return nil
}
