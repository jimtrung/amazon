package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/jimtrung/amazon/internal/logger"
	"github.com/jimtrung/amazon/internal/models"
	"github.com/jimtrung/amazon/internal/services"
)

// GetCart godoc
//	@Summary		Show cart items
//	@Description	Show all items in cart
//	@Tags			Cart
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		models.CartItem		"List of cart items"
//	@Failure		400	{object}	map[string]string	"Bad request error"
//	@Failure		500	{object}	map[string]string	"Internal server error"
//	@Router			/api/v1/cart [get]
func GetCart(c *gin.Context) {
	cart, err := services.GetAllCart()
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
        zap.Any("cart", cart),
        zap.String("url", c.Request.URL.String()),
    )
    defer logger.CloseLogger()

    c.JSON(http.StatusOK, cart)
}

// AddToCart godoc
//	@Summary		Add item to cart
//	@Description	Add a product to cart with  _ quantitys
//	@Tags			Cart
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	map[string]string	"Success messsage"
//	@Failure		400	{object}	map[string]string	"Bad request error"
//	@Failure		500	{object}	map[string]string	"Internal server error"
//	@Router			/api/v1/cart/add [post]
func AddToCart(c *gin.Context) {
    var cartItem models.CartItem

	if err := c.Bind(&cartItem); err != nil {
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

	if err := services.AddToCart(
		cartItem.ProductId,
		cartItem.Quantity,
	); err != nil {
        if err := logger.InitLogger("client/error.log"); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": err.Error(),
            })
            return
        }
        logger.Logger.Error(
            err.Error(),
            zap.String("url", c.Request.URL.String()),
            zap.String("product_id", cartItem.ProductId),
            zap.Int("quantity", cartItem.Quantity),
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
        "Item added to cart",
        zap.String("url", c.Request.URL.String()),
        zap.String("product_id", cartItem.ProductId),
        zap.Int("quantity", cartItem.Quantity),
    )
    defer logger.CloseLogger()

    c.JSON(http.StatusOK, gin.H{"message": "Item added to cart"})
}

// UpdateCart godoc
//	@Summary		Update cart item
//	@Description	Either change the quantity or delete item from cart
//	@Tags			Cart
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	map[string]string	"Success messsage"
//	@Failure		400	{object}	map[string]string	"Bad request error"
//	@Failure		500	{object}	map[string]string	"Internal server error"
//	@Router			/api/v1/cart/update [patch]
func UpdateCart(c *gin.Context) {
    var cartItem models.CartItem
    if err := c.Bind(&cartItem); err != nil {
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

    if err := services.UpdateCartItemQuantity(
        cartItem.ProductId, cartItem.Quantity,
    ); err != nil {
        if err := logger.InitLogger("server/error.log"); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": err.Error(),
            })
            return
        }
        logger.Logger.Info(
            err.Error(),
            zap.String("url", c.Request.URL.String()),
            zap.String("product_id", cartItem.ProductId),
            zap.Int("quantity", cartItem.Quantity),
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
        "Cart updated",
        zap.String("url", c.Request.URL.String()),
        zap.String("product_id", cartItem.ProductId),
        zap.Int("quantity", cartItem.Quantity),
    )
    defer logger.CloseLogger()

    c.JSON(http.StatusOK, gin.H{"message": "Cart updated"})
}

// DeleteFromCart godoc
//	@Summary		Delete item
//	@Description	Remove an item from cart
//	@Tags			Cart
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	map[string]string	"Success messsage"
//	@Failure		400	{object}	map[string]string	"Bad request error"
//	@Failure		500	{object}	map[string]string	"Internal server error"
//	@Router			/api/v1/cart/delete/{productId} [delete]
func DeleteFromCart(c *gin.Context) {
	productId := c.Param("product_id")

    if err := services.DeleteFromCart(productId); err != nil {
        if err := logger.InitLogger("server/error.log"); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": err.Error(),
            })
            return
        }
        logger.Logger.Error(
            err.Error(),
            zap.String("url", c.Request.URL.String()),
            zap.String("product_id", productId),
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
        "Item deleted from cart",
        zap.String("url", c.Request.URL.String()),
        zap.String("product_id", productId),
    )
    defer logger.CloseLogger()

    c.JSON(http.StatusOK, gin.H{"message": "Item deleted from cart"})
}

// DeleteFromCart godoc
//	@Summary		Delete cart database
//	@Description	Delete cart table from database
//	@Tags			Cart
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	map[string]string	"Success messsage"
//	@Failure		400	{object}	map[string]string	"Bad request error"
//	@Failure		500	{object}	map[string]string	"Internal server error"
//	@Router			/protected/drop-cart [delete]
func DropCart(c *gin.Context) {
    if err := services.DropCart(); err != nil {
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
        "Table dropped sucessfully",
        zap.String("url", c.Request.URL.String()),
    )
    defer logger.CloseLogger()

    c.JSON(http.StatusOK, gin.H{"message": "Table dropped successfully"})
}
