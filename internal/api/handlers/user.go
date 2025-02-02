package handlers

import (
	"context"
	"errors"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jimtrung/amazon/internal/api/models"
	"github.com/jimtrung/amazon/internal/config"
)

func GetUsers(c *gin.Context) {
	rows, err := config.DB.Query(context.Background(), "SELECT * FROM users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(
			&user.Id,
			&user.Username,
			&user.Password,
			&user.Email,
			&user.Phone,
			&user.Country,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		users = append(users, user)
	}

	c.JSON(http.StatusOK, users)
}

func AddUser(c *gin.Context) {
	var user models.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	// isValidUserName
	username, err := isValidUsername(user.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := isValidPassword(user.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = config.DB.Exec(
		context.Background(),
		`INSERT INTO users (username, password, email, phone, country) 
		VALUES ($1, $2, $3, $4, $5)`,
		username, user.Password,
		user.Email, user.Phone,
		user.Country,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User added successfully"})
}

func DeleteUser(c *gin.Context) {
	userID := c.Param("user_id")

	_, err := config.DB.Exec(
		context.Background(),
		`DELETE FROM users WHERE id = $1`,
		userID,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func DropUsers(c *gin.Context) {
	dropTable := `
		DROP TABLE users; 
	`

	_, err := config.DB.Exec(
		context.Background(),
		dropTable,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Table dropped successfully"})
}

func isValidUsername(rawUsername string) (string, error) {
	username := strings.ToLower(rawUsername)
	usernameRegex := `^[a-z][a-z0-9._]{2,30}[a-z0-9]$`
	re := regexp.MustCompile(usernameRegex)

	if !re.MatchString(username) {
		return "", errors.New("username must be between 3 and 32 characters, start and end with a letter, and only contain letters, numbers, '.', and '_'")
	}

	return username, nil
}

func isValidPassword(password string) error {
	passwordRegex := `^[A-Za-z\d!@#$%^&*(),.?":{}|<>]{8,64}$`
	re := regexp.MustCompile(passwordRegex)

	if !re.MatchString(password) {
		return errors.New("password must be between 8 and 64 characters and include at least one uppercase letter, one lowercase letter, one number, and one special character")
	}

	return nil
}
