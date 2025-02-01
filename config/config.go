package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

var DB *pgx.Conn
var PORT string

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbURL := loadDBURL()

	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		log.Fatal("Cannot connect to the database:", err)
	}

	DB = conn
	fmt.Println("Database connected successfully")

	PORT = os.Getenv("PORT")
}

func CloseDBConnection() {
	DB.Close(context.Background())
}

func loadDBURL() string {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dbURL := fmt.Sprintf(
		"%s://%s:%s@localhost:%s/%s",
		dbHost,
		dbUser,
		dbPass,
		dbPort,
		dbName,
	)

	return dbURL
}
