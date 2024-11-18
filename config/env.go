package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() (string, string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	login := os.Getenv("LOGIN")
	password := os.Getenv("PASSWORD")

	if login == "" || password == "" {
		log.Fatalf("The LOGIN and PASSWORD environment variables are required.")
	}

	return login, password
}
