package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jimtrung/amazon/config"
	"github.com/jimtrung/amazon/internal/models"
	"github.com/jimtrung/amazon/internal/services"
	"github.com/jimtrung/amazon/internal/logger"

	"golang.org/x/crypto/bcrypt"
	"go.uber.org/zap"
)

// GetUsers godoc
//	@Summary		Show all the users
//	@Description	Show all the users and infos
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		models.User			"List of users"
//	@Failure		400	{object}	map[string]string	"Bad request error"
//	@Failure		500	{object}	map[string]string	"Internal server error"
//	@Router			/api/v1/users [get]
func GetUsers(c *gin.Context) {
	users, err := services.GetUsers()
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
        "Successfully get users",
        zap.String("url", c.Request.URL.String()),
        zap.Any("users", users),
    )
    defer logger.CloseLogger()

	c.JSON(http.StatusOK, users)
}

// Signup godoc
//	@Summary		Create a new user
//	@Description	Create a user if given info is valid
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	map[string]string	"Success message"
//	@Failure		400	{object}	map[string]string	"Bad request error"
//	@Failure		500	{object}	map[string]string	"Internal server error"
//	@Router			/api/v1/users/signup [post]
func Signup(c *gin.Context) {
	var user models.User
	if err := c.Bind(&user); err != nil {
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

		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	username, hash, err := services.IsValidUser(user)
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

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.AddUser(username, hash, user); err != nil {
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
        "User added successfully",
        zap.String("url", c.Request.URL.String()),
        zap.String("username", username),
    )
    defer logger.CloseLogger()

	c.JSON(http.StatusOK, gin.H{"message": "User added successfully"})
}

// DeleteUser godoc
//	@Summary		Delete a user with a given id
//	@Description	Delete user with id in the URL path
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	map[string]string	"Success message"
//	@Failure		400	{object}	map[string]string	"Bad request error"
//	@Failure		500	{object}	map[string]string	"Internal server error"
//	@Router			/protected/delete/{user_id} [delete]
func DeleteUser(c *gin.Context) {
	userID := c.Param("user_id")

	if err := services.DeleteUser(userID); err != nil {
        if err := logger.InitLogger("server/error.log"); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": err.Error(),
            })
            return
        }
        logger.Logger.Error(
            err.Error(),
            zap.String("url", c.Request.URL.String()),
            zap.String("user_id", userID),
        )
        defer logger.CloseLogger()

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
    if err := logger.InitLogger("server/action.log"); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }
    logger.Logger.Info(
        "User deleted successfully",
        zap.String("url", c.Request.URL.String()),
    )
    defer logger.CloseLogger()

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

// DropUsers godoc
//	@Summary		Drop users table
//	@Description	Drop users table in the database
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	map[string]string	"Success message"
//	@Failure		400	{object}	map[string]string	"Bad request error"
//	@Failure		500	{object}	map[string]string	"Internal server error"
//	@Router			/protected/drop-users [delete]
func DropUsers(c *gin.Context) {
	if err := services.DropUser(); err != nil {
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

// Login godoc
//	@Summary		Login to an existed account
//	@Description	Login to an account with validation
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	map[string]string	"Success message"
//	@Failure		400	{object}	map[string]string	"Bad request error"
//	@Failure		500	{object}	map[string]string	"Internal server error"
//	@Router			/api/v1/users/login [post]
func Login(c *gin.Context) {
	var user models.UserResponse
	if err := c.Bind(&user); err != nil {
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

	row := config.DB.QueryRow(
		context.Background(),
		`SELECT password FROM users WHERE username = $1`,
		user.Username,
	)

	var hashedPassword string
	err := row.Scan(&hashedPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Database error",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(user.Password),
	); err != nil {
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

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
    if err := logger.InitLogger("server/action.log"); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }
    logger.Logger.Info(
        "Login successfully",
        zap.String("url", c.Request.URL.String()),
        zap.String("username", user.Username),
    )
    defer logger.CloseLogger()

	c.JSON(http.StatusOK, gin.H{"message": "Login successfully"})
}
