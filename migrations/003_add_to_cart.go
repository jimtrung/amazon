package migrations

import (
	"context"
	"fmt"

	"github.com/jimtrung/amazon/config"
)

func CreateAddToCart() error {
	addToCart := `
		CREATE OR REPLACE FUNCTION add_to_cart(add_id VARCHAR(255), add_quantity INT) 
		RETURNS VOID AS $$
		BEGIN
			IF EXISTS (SELECT 1 FROM cart WHERE product_id = add_id) THEN
				UPDATE cart
				SET quantity = quantity + add_quantity
				WHERE product_id = add_id;
			ELSE 
				INSERT INTO cart VALUES (add_id, add_quantity);
			END IF;
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

	fmt.Println("003_add_to_cart(1/1) - Add to cart function created successfully")
	return nil
}
