package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jimtrung/amazon/internal/models"
	"github.com/jimtrung/amazon/internal/services"
)

func GetProducts(c *gin.Context) {
	products, err := services.GetProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, products)
}

func DropProducts(c *gin.Context) {
	if err := services.DropCart(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
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

	if err := services.Transfer(products); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transfer successfully"})
}
