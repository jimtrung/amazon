package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jimtrung/amazon/internal/logger"
	"github.com/jimtrung/amazon/internal/models"
	"github.com/jimtrung/amazon/internal/services"
)

// GetProducts godoc
//
//	@Summary		Show products
//	@Description	Show all the products of the website
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		models.Product		"List of products"
//	@Failure		400	{object}	map[string]string	"Bad request error"
//	@Failure		500	{object}	map[string]string	"Internal server error"
//	@Router			/api/v1/product [get]
func GetProducts(c *gin.Context) {
	products, err := services.GetProducts()
	if err != nil {
		logger.LogAndRespond(
			c, "server/error.log", "Failed to get products from database",
			err, http.StatusInternalServerError,
		)
		return
	}

	logger.LogAndRespond(
		c, "server/action.log", "Successfully get products",
		nil, http.StatusOK, products,
	)
}

// Transfer godoc
//
//	@Summary		Insert products to table
//	@Description	Insert a JSON of products to table
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	map[string]string	"Success message"
//	@Failure		400	{object}	map[string]string	"Bad request error"
//	@Failure		500	{object}	map[string]string	"Internal server error"
//	@Router			/api/v1/products/transfer [post]
func Transfer(c *gin.Context) {
	var products []models.Product

	if err := c.Bind(&products); err != nil {
		logger.LogAndRespond(
			c, "server/error.log", "Wrong JSON format",
			err, http.StatusBadRequest,
		)
		return
	}

	if err := services.Transfer(products); err != nil {
		logger.LogAndRespond(
			c, "server/error.log", "Failed to transfer products to database",
			err, http.StatusInternalServerError,
		)
		return
	}

	logger.LogAndRespond(
		c, "server/action.log", "Products data transfered successfully",
		nil, http.StatusOK,
	)
}
