package auth

import (
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	username string
	password string
}

var users = []User{
	{username: "trung123", password: "int"},
	{username: "dajknsd", password: "hafjk"},
	{username: "dabd", password: "asd"},
	{username: "driqeq", password: "int"},
	{username: "dabmd", password: "int"},
}

func decodeBase64(authorizationData string) (string, string, error) {
	clientEncoded := authorizationData[6:]
	decodedBytes, err := base64.StdEncoding.DecodeString(clientEncoded)
	if err != nil {
		return "", "", err
	}

	decodedString := string(decodedBytes)
	var username, password string
	isAfterColon := false

	for _, char := range decodedString {
		if !isAfterColon {
			if char == ':' {
				isAfterColon = true
			} else {
				username += string(char)
			}
		} else {
			password += string(char)
		}
	}
	return username, password, nil
}

func isValidUser(clUser, clPass string) bool {
	for _, user := range users {
		if user.username == clUser && user.password == clPass {
			return true
		}
	}
	return false
}

func BasicAuth(c *gin.Context) {
	authorization := c.GetHeader("Authorization")
	if authorization == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing header",
		})
		return
	}

	username, password, err := decodeBase64(authorization)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if isValidUser(username, password) {
		c.JSON(http.StatusOK, gin.H{"message": "Login successfully!"})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Wrong username/password!"})
}
