package migrations

import (
	"context"
	"fmt"

	"github.com/jimtrung/amazon/internal/config"
)

func CreateUpdateCart() error {
	addToCart := `
		CREATE OR REPLACE FUNCTION update_cart(update_id VARCHAR(255), new_quantity INT) 
		RETURNS VOID AS $$
		BEGIN
			IF new_quantity <= 0 THEN
				DELETE FROM cart WHERE product_id = update_id;
			ELSE
				UPDATE cart
				SET quantity = new_quantity
				WHERE product_id = update_id;
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

	fmt.Println("005_update_cart(1/1) - Update cart function created successfully")
	return nil
}
