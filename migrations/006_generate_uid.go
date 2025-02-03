package migrations

import (
	"context"
	"fmt"

	"github.com/jimtrung/amazon/config"
)

func CreateUIDGenerator() error {
	generateUID := `
		CREATE OR REPLACE FUNCTION generate_uid()
		RETURNS BIGINT AS $$
		DECLARE
			new_id BIGINT;
		BEGIN
			LOOP
				new_id := 1000000000 + floor(random() * 9000000000)::BIGINT;

				EXIT WHEN NOT EXISTS (SELECT 1 FROM users WHERE id = new_id);
			END LOOP;
			RETURN new_id;
		END;
		$$ LANGUAGE plpgsql;
	`

	_, err := config.DB.Exec(
		context.Background(),
		generateUID,
	)
	if err != nil {
		return err
	}

	fmt.Println("006_generate_uid(1/1) - Generate UID function created successfully")
	return nil
}
