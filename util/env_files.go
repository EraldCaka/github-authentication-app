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
	ACTIVE_TOKEN  string
)

func InitEnvironmentVariables() {
	err := godotenv.Load(".env", "token.env")
	if err != nil {
		log.Fatal(err)
	}

	DB_URL = os.Getenv("DB_URL")
	CLIENT_ID = os.Getenv("CLIENT_ID")
	CLIENT_SECRET = os.Getenv("CLIENT_SECRET")
	ACTIVE_TOKEN = os.Getenv("ACTIVE_TOKEN")
}

func SetActiveToken(activeToken string) error {
	err := godotenv.Write(map[string]string{
		"ACTIVE_TOKEN": activeToken,
	}, "token.env")
	if err != nil {
		return err
	}
	return nil
}
func ClearActiveToken() error {
	err := godotenv.Write(map[string]string{
		"ACTIVE_TOKEN": "",
	}, "token.env")
	if err != nil {
		return err
	}
	return nil
}
func init() {
	err := godotenv.Load("token.env")
	if err != nil {
		log.Fatal(err)
	}
	ACTIVE_TOKEN = os.Getenv("ACTIVE_TOKEN")
}
