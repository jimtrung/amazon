package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jimtrung/amazon/config"
	"github.com/jimtrung/amazon/internal/models"
)

func GetProducts(c *gin.Context) {
	rows, err := config.DB.Query(context.Background(), "SELECT * FROM products")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Database query failed: " + err.Error(),
		})
		return
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Image,
			&product.Rating.Stars,
			&product.Rating.Count,
			&product.PriceCents,
			&product.Keywords,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		products = append(products, product)
	}

	c.JSON(http.StatusOK, products)
}

func DropProducts(c *gin.Context) {
	dropTable := `
		DROP TABLE products;
	`

	_, err := config.DB.Exec(
		context.Background(),
		dropTable,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Table dropped successfully"})
}

func Transfer(c *gin.Context) {
	var products []models.Product

	if err := c.Bind(&products); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	for _, product := range products {
		_, err := config.DB.Exec(
			context.Background(),
			"INSERT INTO products VALUES ($1, $2, $3, $4, $5, $6, $7)",
			product.Id, product.Name, product.Image,
			product.Rating.Stars, product.Rating.Count,
			product.PriceCents, product.Keywords,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transfer successfully"})
}
