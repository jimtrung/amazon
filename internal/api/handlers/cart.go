package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jimtrung/amazon/config"
	"github.com/jimtrung/amazon/models"
)

func GetCart(c *gin.Context) {
	rows, err := config.DB.Query(context.Background(), "SELECT * FROM cart")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer rows.Close()

	var cart []models.CartItem
	for rows.Next() {
		var cartItem models.CartItem
		err := rows.Scan(
			&cartItem.ProductId,
			&cartItem.Quantity,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		cart = append(cart, cartItem)
	}

	c.JSON(http.StatusOK, cart)
}

// Add item to cart
func AddToCart(c *gin.Context) {
	var cartItem models.CartItem
	if err := c.Bind(&cartItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err := config.DB.Exec(
		context.Background(),
		"SELECT add_to_cart($1, $2);",
		cartItem.ProductId, cartItem.Quantity,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item added to cart"})
}

func UpdateCart(c *gin.Context) {
	var cartItem models.CartItem
	if err := c.Bind(&cartItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err := config.DB.Exec(
		context.Background(),
		"SELECT update_cart($1, $2)",
		cartItem.ProductId, cartItem.Quantity,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cart updated"})
}

func DeleteFromCart(c *gin.Context) {
	var cartItem models.CartItem
	if err := c.Bind(&cartItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err := config.DB.Exec(
		context.Background(),
		"SELECT delete_from_cart($1)",
		cartItem.ProductId,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item deleted from cart"})
}

func DropCart(c *gin.Context) {
	dropTable := `
		DROP TABLE cart; 
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
