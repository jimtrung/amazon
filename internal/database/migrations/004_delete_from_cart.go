package migrations

import (
	"context"
	"fmt"

	"github.com/jimtrung/amazon/config"
)

func CreateDeleteFromCart() error {
	deleteFromCart := `
		CREATE OR REPLACE FUNCTION delete_from_cart(delete_id VARCHAR(255)) 
		RETURNS VOID AS $$
		BEGIN
			IF EXISTS (SELECT 1 FROM cart WHERE product_id = delete_id) THEN
				DELETE FROM cart WHERE product_id = delete_id;
			ELSE
				RAISE EXCEPTION 'Can not find item in cart';
			END IF;
		END;
		$$ LANGUAGE plpgsql;
	`

	_, err := config.DB.Exec(
		context.Background(),
		deleteFromCart,
	)
	if err != nil {
		return err
	}

	fmt.Println("004_delete_from_cart(1/1) - Delete from cart function created successfully")
	return nil
}
