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
        logger.LogAndRespond(
            c, "server/error.log", "Failed to get users from database",
            err, http.StatusInternalServerError,
        )
        return
	}

    logger.LogAndRespond(
        c, "server/action.log", "Successfully get users",
        nil, http.StatusOK, users,
    )
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
        logger.LogAndRespond(
            c, "server/error.log", "Wrong JSON format",
            err, http.StatusBadRequest,
        )
        return
	}

	username, hash, err := services.IsValidUser(user)
	if err != nil {
        logger.LogAndRespond(
            c, "client/error.log", "Invalid user format",
            err, http.StatusBadRequest,
        )
        return
	}

	if err := services.AddUser(username, hash, user); err != nil {
        logger.LogAndRespond(
            c, "server/error.log", "Failed to add user to database",
            err, http.StatusInternalServerError,
        )
        return
	}

    logger.LogAndRespond(
        c, "client/action.log", "User created successfully",
        nil, http.StatusOK,
    )
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
        logger.LogAndRespond(
            c, "server/error.log", "Failed to delete user from database",
            err, http.StatusInternalServerError,
        )
        return
	}

    logger.LogAndRespond(
        c, "client/action.log", "User deleted successfully",
        nil, http.StatusOK,
    )
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
        logger.LogAndRespond(
            c, "server/error.log", "Wrong JSON format",
            err, http.StatusBadRequest,
        )
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
        logger.LogAndRespond(
            c, "client/error.log", "No user found",
            err, http.StatusBadRequest,
        )
        return
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(user.Password),
	); err != nil {
        logger.LogAndRespond(
            c, "client/action.log", "Wrong password",
            nil, http.StatusOK, user.Username,
        )
        return
    }

    logger.LogAndRespond(
        c, "client/action.log", "Login successfully",
        err, http.StatusOK,
    )
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

