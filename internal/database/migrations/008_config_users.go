package migrations

import (
	"context"
	"fmt"

	"github.com/jimtrung/amazon/internal/config"
)

func BindDefaultUsers() error {
	usersDefault := `
		ALTER TABLE users
		ALTER COLUMN id SET DEFAULT generate_uid(),
		ALTER COLUMN email SET DEFAULT 'no_email',
		ALTER COLUMN phone SET DEFAULT 'no_phone',
		ALTER COLUMN country SET DEFAULT 'global';
	`

	_, err := config.DB.Exec(
		context.Background(),
		usersDefault,
	)
	if err != nil {
		return err
	}

	fmt.Println("008_config_users(1/2) - Bind default users successfully")
	return nil
}

func BindRuleUsers() error {
	usersRule := `
		-- No rule for now
	`

	_, err := config.DB.Exec(
		context.Background(),
		usersRule,
	)
	if err != nil {
		return err
	}

	fmt.Println("008_config_users(2/2) - Bind rules users successfully")
	return nil
}
