package util

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	SECRET_KEY    string
	DB_URL        string
	CLIENT_ID     string
	CLIENT_SECRET string
)

func InitEnvironmentVariables() {
	err := godotenv.Load(".env", "client.env")
	if err != nil {
		log.Fatal(err)
	}

	DB_URL = os.Getenv("DB_URL")
	SECRET_KEY = os.Getenv("SECRET_KEY")
	CLIENT_ID = os.Getenv("CLIENT_ID")
	CLIENT_SECRET = os.Getenv("client_secret")
}
