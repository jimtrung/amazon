package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jimtrung/amazon/internal/models"
	"github.com/jimtrung/amazon/internal/services"
)

func GetCart(c *gin.Context) {
	cart, err := services.GetAllCart()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, cart)
}

func AddToCart(c *gin.Context) {
	var cartItem models.CartItem
	if err := c.Bind(&cartItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := services.AddToCart(
		cartItem.ProductId,
		cartItem.Quantity,
	); err != nil {
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

    if err := services.UpdateCartItemQuantity(
        cartItem.ProductId, cartItem.Quantity,
    ); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
    }

	c.JSON(http.StatusOK, gin.H{"message": "Cart updated"})
}

func DeleteFromCart(c *gin.Context) {
	productId := c.Param("product_id")

    if err := services.DeleteFromCart(productId); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
    }

	c.JSON(http.StatusOK, gin.H{"message": "Item deleted from cart"})
}

func DropCart(c *gin.Context) {
    if err := services.DropCart(); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
    }

	c.JSON(http.StatusOK, gin.H{"message": "Table dropped successfully"})
}
