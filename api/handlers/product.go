package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jimtrung/amazon/internal/logger"
	"github.com/jimtrung/amazon/internal/models"
	"github.com/jimtrung/amazon/internal/services"
	"go.uber.org/zap"
)

func GetProducts(c *gin.Context) {
	products, err := services.GetProducts()
	if err != nil {
        if err := logger.InitLogger("client/error.log"); err != nil {
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

func Transfer(c *gin.Context) {
	var products []models.Product

	if err := c.Bind(&products); err != nil {
        if err := logger.InitLogger("client/error.log"); err != nil {
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
        if err := logger.InitLogger("client/error.log"); err != nil {
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
