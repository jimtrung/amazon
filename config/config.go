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

	dbURL := os.Getenv("DB_URL")

	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		log.Fatal("Cannot connect to the database:", err)
	}

	DB = conn
	PORT = os.Getenv("PORT")
	fmt.Println("Database connected successfully")
}
