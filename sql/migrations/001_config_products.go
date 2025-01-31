package migrations

import (
	"context"
	"fmt"

	"github.com/jimtrung/amazon/config"
)

func BindDefaultProducts() error {
	productsDefault := `
		ALTER TABLE products
		ALTER COLUMN image SET DEFAULT 'no_path',
		ALTER COLUMN name SET DEFAULT 'no_name',
		ALTER COLUMN rating_stars SET DEFAULT 0.0,
		ALTER COLUMN rating_count SET DEFAULT 0,
		ALTER COLUMN price_cents SET DEFAULT 0,
		ALTER COLUMN keywords SET DEFAULT ARRAY[]::TEXT[];
	`

	_, err := config.DB.Exec(
		context.Background(),
		productsDefault,
	)
	if err != nil {
		return err
	}

	fmt.Println("001_config_products - Bind default products successfully")
	return nil
}

func BindRuleProducts() error {
	productsRule := `
		DO $$
		BEGIN
			-- Check and add constraint for rating_stars range if not exists
			IF NOT EXISTS (
				SELECT 1 FROM pg_constraint
				WHERE conname = 'check_products_rating_stars_range'
			) THEN
				ALTER TABLE products
				ADD CONSTRAINT check_products_rating_stars_range
				CHECK (0 <= rating_stars AND rating_stars <= 5.0);
			END IF;

			-- Check and add constraint for rating_count if not exists
			IF NOT EXISTS (
				SELECT 1 FROM pg_constraint
				WHERE conname = 'check_products_rating_count'
			) THEN
				ALTER TABLE products
				ADD CONSTRAINT check_products_rating_count
				CHECK (rating_count >= 0);
			END IF;

			-- Check and add constraint for price_cents if not exists
			IF NOT EXISTS (
				SELECT 1 FROM pg_constraint
				WHERE conname = 'check_products_price_cents'
			) THEN
				ALTER TABLE products
				ADD CONSTRAINT check_products_price_cents
				CHECK (price_cents >= 0);
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

	fmt.Println("001_config_products - Bind rules products successfully")
	return nil
}
