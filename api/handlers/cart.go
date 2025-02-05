package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

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
        logger.LogAndRespond(
            c, "server/error.log", "Failed to get cart from database",
            err, http.StatusInternalServerError,
        )
        return
    }

    logger.LogAndRespond(
        c, "server/action.log", "Successfully get cart",
        nil, http.StatusOK, cart,
    )
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
        logger.LogAndRespond(
            c, "server/error.log", "Wrong JSON format",
            err, http.StatusBadRequest,
        )
        return
	}

	if err := services.AddToCart(
		cartItem.ProductId,
		cartItem.Quantity,
    ); err != nil {
        logger.LogAndRespond(
            c, "server/error.log", "Failed to add item cart",
            err, http.StatusInternalServerError,
        )
        return
	}

    logger.LogAndRespond(
        c, "client/action.log", "Item added to cart",
        nil, http.StatusOK, cartItem,
    )
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
        logger.LogAndRespond(
            c, "server/error.log", "Wrong JSON format",
            err, http.StatusBadRequest,
        )
        return
    }

    if err := services.UpdateCartItemQuantity(
        cartItem.ProductId, cartItem.Quantity,
    ); err != nil {
        logger.LogAndRespond(
            c, "server/error.log", "Failed to update item quantity",
            err, http.StatusInternalServerError,
        )
        return
    }

    logger.LogAndRespond(
        c, "client/action.log", "Item quantity updated successfully",
        nil, http.StatusOK, cartItem,
    )
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
        logger.LogAndRespond(
            c, "server/error.log", "product_id not found",
            err, http.StatusInternalServerError,
        )
        return
    }
    logger.LogAndRespond(
        c, "client/action.log", "Item deleted successfully",
        nil, http.StatusOK, productId,
    )
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
        logger.LogAndRespond(
            c, "server/error.log", "Failed to drop table",
            err, http.StatusInternalServerError,
        )
        return
    }

    logger.LogAndRespond(
        c, "server/action.log", "Table dropped successfully",
        nil, http.StatusOK,
    )
}
