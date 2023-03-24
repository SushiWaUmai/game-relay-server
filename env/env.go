package env

import (
	"log"

	"github.com/joho/godotenv"
)

func init() {
	setupDotenv()
}

func setupDotenv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	loadEnv()
}
