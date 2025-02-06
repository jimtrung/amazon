package services

import (
	"context"

	"github.com/jimtrung/amazon/config"
	"github.com/jimtrung/amazon/internal/models"
)

func GetProducts() ([]models.Product, error) {
	rows, err := config.DB.Query(context.Background(), "SELECT * FROM products")
	if err != nil {
		return []models.Product{}, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		err := rows.Scan(
			&product.Id, &product.Name, &product.Image,
			&product.Rating.Stars, &product.Rating.Count,
			&product.PriceCents, &product.Keywords,
			&product.CreatedAt, &product.UpdatedAt,
		)
		if err != nil {
			return []models.Product{}, err
		}
		products = append(products, product)
	}
	return products, nil
}

func Transfer(products []models.Product) error {
	for _, product := range products {
		_, err := config.DB.Exec(
			context.Background(),
			"INSERT INTO products VALUES ($1, $2, $3, $4, $5, $6, $7)",
			product.Id, product.Name, product.Image,
			product.Rating.Stars, product.Rating.Count,
			product.PriceCents, product.Keywords,
		)
		if err != nil {
			return err
		}
	}
	return nil
}
