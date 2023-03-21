package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(envName string) string {
	if os.Getenv("GIN_MODE") != "release" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	return os.Getenv(envName)
}
