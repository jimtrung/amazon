package auth

import (
	"encoding/base64"

	"github.com/gofiber/fiber/v3"
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

func BasicAuth(c fiber.Ctx) error {
	authorization := c.Get("Authorization")
	if authorization == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Missing Authorization"})
	}

	username, password, err := decodeBase64(authorization)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Failed to decode"})
	}

	if isValidUser(username, password) {
		return c.JSON(fiber.Map{"message": "Login successfully!"})
	}

	return c.Status(401).JSON(fiber.Map{"error": "Wrong username/password"})
}
