package migrations

import (
	"context"
	"fmt"

	"github.com/jimtrung/amazon/config"
)

func BindDefaultCart() error {
	cartDefault := `
		ALTER TABLE cart
		ALTER COLUMN quantity SET DEFAULT 0;
	`

	_, err := config.DB.Exec(
		context.Background(),
		cartDefault,
	)
	if err != nil {
		return err
	}

	fmt.Println("002_config_cart(1/2) - Bind default cart successfully")
	return nil
}

func BindRuleCart() error {
	cartRule := `
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
		cartRule,
	)
	if err != nil {
		return err
	}

	fmt.Println("002_config_cart(2/2) - Bind rules cart successfully")
	return nil
}
