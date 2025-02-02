package migrations

import (
	"context"
	"fmt"

	"github.com/jimtrung/amazon/internal/config"
)

func CreateUIDGenerator() error {
	generateUID := `
		CREATE OR REPLACE FUNCTION generate_uid(size INT) RETURNS TEXT AS $$
		DECLARE
			characters TEXT := 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
			bytes BYTEA := gen_random_bytes(size);
			l INT := length(characters);
			i INT := 0;
			output TEXT := '';
		BEGIN
			WHILE i < size LOOP
				output := output || substr(characters, get_byte(bytes, i) % l + 1, 1);
				i := i + 1;
			END LOOP;
			RETURN output;
		END;
		$$ LANGUAGE plpgsql VOLATILE;
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
