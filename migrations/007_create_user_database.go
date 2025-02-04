package migrations

import (
	"context"
	"fmt"

	"github.com/jimtrung/amazon/config"
)

func CreateTableUser() error {
	createUser := `
		CREATE TABLE IF NOT EXISTS users (
			id 		 BIGINT,
			username VARCHAR(255),
			password VARCHAR(255),
			email    VARCHAR(255),
			phone    VARCHAR(255),
			country  VARCHAR(255),
			status   VARCHAR(255),
			PRIMARY KEY (id, username),
			CONSTRAINT unique_id UNIQUE (id),
    		CONSTRAINT unique_username UNIQUE (username)
		);
	`

	_, err := config.DB.Exec(
		context.Background(),
		createUser,
	)
	if err != nil {
		return err
	}

	fmt.Println("007_create_user_database(1/1) - Table user created successfully")
	return nil
}
