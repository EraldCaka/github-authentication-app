package util

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	DB_URL        string
	CLIENT_ID     string
	CLIENT_SECRET string
)

func InitEnvironmentVariables() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	DB_URL = os.Getenv("DB_URL")
	CLIENT_ID = os.Getenv("CLIENT_ID")
	CLIENT_SECRET = os.Getenv("CLIENT_SECRET")
}
