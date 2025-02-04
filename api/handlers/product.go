package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jimtrung/amazon/internal/logger"
	"github.com/jimtrung/amazon/internal/models"
	"github.com/jimtrung/amazon/internal/services"

	"go.uber.org/zap"
)

// GetProducts godoc
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
        if err := logger.InitLogger("server/error.log"); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": err.Error(),
            })
            return
        }
        logger.Logger.Error(
            err.Error(),
            zap.String("url", c.Request.URL.String()),
        )
        defer logger.CloseLogger()

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
    if err := logger.InitLogger("client/action.log"); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }
    logger.Logger.Info(
        "Successfully get products",
        zap.String("url", c.Request.URL.String()),
        zap.Any("products", products),
    )
    defer logger.CloseLogger()

    c.JSON(http.StatusOK, products)
}

// DropProducts godoc
//	@Summary		Delete products table from database
//	@Description	Remove table from database
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	map[string]string	"Success message"
//	@Failure		400	{object}	map[string]string	"Bad request error"
//	@Failure		500	{object}	map[string]string	"Internal server error"
//	@Router			/protected/drop-products [delete]
func DropProducts(c *gin.Context) {
	if err := services.DropProducts(); err != nil {
        if err := logger.InitLogger("server/error.log"); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": err.Error(),
            })
            return
        }
        logger.Logger.Error(
            err.Error(),
            zap.String("url", c.Request.URL.String()),
        )
        defer logger.CloseLogger()

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
        return
	}
    if err := logger.InitLogger("server/action.log"); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }
    logger.Logger.Info(
        "Table dropped successfully",
        zap.String("url", c.Request.URL.String()),
    )
    defer logger.CloseLogger()

	c.JSON(http.StatusOK, gin.H{"message": "Table dropped successfully"})
}

// Transfer godoc
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
        if err := logger.InitLogger("server/error.log"); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": err.Error(),
            })
            return
        }
        logger.Logger.Error(
            err.Error(),
            zap.String("url", c.Request.URL.String()),
        )
        defer logger.CloseLogger()

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := services.Transfer(products); err != nil {
        if err := logger.InitLogger("server/error.log"); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": err.Error(),
            })
            return
        }
        logger.Logger.Error(
            err.Error(),
            zap.String("url", c.Request.URL.String()),
            zap.Any("products", products),
        )
        defer logger.CloseLogger()

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
        return
	}
    if err := logger.InitLogger("client/action.log"); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }
    logger.Logger.Info(
        "Successfully get products",
        zap.String("url", c.Request.URL.String()),
    )
    defer logger.CloseLogger()

	c.JSON(http.StatusOK, gin.H{"message": "Transfer successfully"})
}
