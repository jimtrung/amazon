package sql

import (
	"context"
	"fmt"

	"github.com/jimtrung/amazon/config"
)

func CreateSchema() error {
	schema := `
		CREATE TABLE IF NOT EXISTS products (
			id   		 VARCHAR(255) PRIMARY KEY,
			image		 VARCHAR(255),
			name 		 VARCHAR(255) NOT NULL,
			rating_stars FLOAT,
			rating_count INT,
			price_cents  INT,
			keywords     TEXT[]
		);

		CREATE TABLE IF NOT EXISTS cart (
			product_id VARCHAR(255),
			quantity   INT,
			FOREIGN KEY (product_id) REFERENCES products(id)
		);
	`
	_, err := config.DB.Exec(
		context.Background(),
		schema,
	)
	if err != nil {
		return err
	}

	fmt.Println("schema.go - Create schema successfully")
	return nil
}
