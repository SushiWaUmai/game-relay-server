package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	setupDotenv()
}

func setupDotenv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
		os.Exit(1)
	}

	loadEnv()
}
