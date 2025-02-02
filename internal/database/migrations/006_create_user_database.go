package migrations

import (
	"context"
	"fmt"

	"github.com/jimtrung/amazon/internal/config"
)

func CreateTableUser() error {
	createUser := `
		CREATE TABLE IF NOT EXISTS users (
			id 		 INT PRIMARY KEY,
			username VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL,
			email    VARCHAR(255),
			phone    VARCHAR(255),
			country  VARCHAR(255)
		);
	`

	_, err := config.DB.Exec(
		context.Background(),
		createUser,
	)
	if err != nil {
		return err
	}

	fmt.Println("006_create_user_database(1/1) - Table user created successfully")
	return nil
}
