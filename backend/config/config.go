package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/ilyes-rhdi/buildit-Gql/pkg/logger"
)

var JWT_SECRET string
var EMAIL string
var EMAIL_PASSWORD string

func Load() {
	// Load env file if present
	_ = godotenv.Load()

	JWT_SECRET = os.Getenv("JWT_SECRET")
	EMAIL = os.Getenv("EMAIL")
	EMAIL_PASSWORD = os.Getenv("EMAIL_PASSWORD")

	// Initialize the logger
	logger.NewLogger()
}
