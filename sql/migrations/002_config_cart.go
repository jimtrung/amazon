package migrations

import (
	"context"
	"fmt"

	"github.com/jimtrung/amazon/config"
)

func BindDefaultCart() error {
	productsDefault := `
		ALTER TABLE cart
		ALTER COLUMN quantity SET DEFAULT 0;
	`

	_, err := config.DB.Exec(
		context.Background(),
		productsDefault,
	)
	if err != nil {
		return err
	}

	fmt.Println("002_config_cart - Bind default cart successfully")
	return nil
}

func BindRuleCart() error {
	productsRule := `
		-- Make sure quantity >= 0
		DO $$
		BEGIN
			-- Check and add constraint for cart quantity
			IF NOT EXISTS (
				SELECT 1 FROM pg_constraint
				WHERE conname = 'check_cart_quantity'
			) THEN
				ALTER TABLE cart
				ADD CONSTRAINT check_cart_quantity
				CHECK (quantity >= 0);
			END IF;
		END $$;
	`

	_, err := config.DB.Exec(
		context.Background(),
		productsRule,
	)
	if err != nil {
		return err
	}

	fmt.Println("002_config_cart - Bind rules cart successfully")
	return nil
}
