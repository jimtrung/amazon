package services

import (
	"context"
	"errors"
	"regexp"
	"strings"

	"github.com/jimtrung/amazon/config"
	"github.com/jimtrung/amazon/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers() ([]models.User, error) {
	rows, err := config.DB.Query(context.Background(), "SELECT * FROM users")
	if err != nil {
		return []models.User{}, err
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
			return []models.User{}, err
		}
		users = append(users, user)
	}
	return users, nil
}

func IsValidUser(user models.User) (string, []byte, error) {
	// isValidUserName
	username, err := IsValidUsername(user.Username)
	if err != nil {
		return "", nil, err
	}

	if err := IsValidPassword(user.Password); err != nil {
		return "", nil, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return "", nil, err
	}
	return username, hash, nil
}

func AddUser(username string, hash []byte, user models.User) error {
	_, err := config.DB.Exec(
		context.Background(),
		`INSERT INTO users (username, password, email, phone, country) 
		VALUES ($1, $2, $3, $4, $5)`,
		username, hash,
		user.Email, user.Phone,
		user.Country,
	)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(userID string) error {
	_, err := config.DB.Exec(
		context.Background(),
		`DELETE FROM users WHERE id = $1`,
		userID,
	)
	if err != nil {
		return err
	}
	return nil
}

func DropUser() error {
	dropTable := `
		DROP TABLE users; 
	`

	_, err := config.DB.Exec(
		context.Background(),
		dropTable,
	)
	if err != nil {
		return err
	}
	return nil
}

func IsValidUsername(rawUsername string) (string, error) {
	username := strings.ToLower(rawUsername)
	usernameRegex := `^[a-z][a-z0-9._]{2,30}[a-z0-9]$`
	re := regexp.MustCompile(usernameRegex)

	if !re.MatchString(username) {
		return "", errors.New(`username must be between 3 and 32 characters,
start and end with a letter, and only contain letters, numbers, '.', and '_'`)
	}

	return username, nil
}

func IsValidPassword(password string) error {
	passwordRegex := `^[A-Za-z\d!@#$%^&*(),.?":{}|<>]{8,64}$`
	re := regexp.MustCompile(passwordRegex)

	if !re.MatchString(password) {
		return errors.New(`password must be between 8 and 64 characters and 
include at least one uppercase letter, one lowercase letter, one number, and 
one special character`)
	}

	return nil
}
