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
		ALTER TABLE cart
		ADD CONSTRAINT check_quantity
		CHECK (quantity >= 0);
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
