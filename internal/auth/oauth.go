package auth

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

const (
	key    = "trung123"
	maxAge = 86400 * 30
	isProd = false
)

func NewAuth() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	googleClientId := os.Getenv("GOOGLE_CLIENT_ID")
	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/auth/oauth"
	store.Options.HttpOnly = true
	store.Options.Secure = isProd

	gothic.Store = store

	goth.UseProviders(
		google.New(
			googleClientId,
			googleClientSecret,
			"http://127.0.0.1:8008/auth/oauth",
		),
	)
}

func GetAuthCallbackFunction(c fiber.Ctx) error {
	user, err := gothic.CompleteUserAuth(c)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": err})
	}

	return c.JSON(fiber.Map{"message": user})
}
