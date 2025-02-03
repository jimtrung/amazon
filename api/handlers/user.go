package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jimtrung/amazon/config"
	"github.com/jimtrung/amazon/internal/models"
	"github.com/jimtrung/amazon/internal/services"

	"golang.org/x/crypto/bcrypt"
)

func GetUsers(c *gin.Context) {
	users, err := services.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, users)
}

func Signup(c *gin.Context) {
	var user models.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	username, hash, err := services.IsValidUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.AddUser(username, hash, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{"message": "User added successfully"})
}

func DeleteUser(c *gin.Context) {
	userID := c.Param("user_id")

	if err := services.DeleteUser(userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func DropUsers(c *gin.Context) {
	if err := services.DropUser(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Table dropped successfully"})
}

func Login(c *gin.Context) {
	var user models.UserResponse
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successfully"})
}
