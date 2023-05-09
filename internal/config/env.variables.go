package config

import (
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/exp/slog"
)

// GetEnv returns an environment variable or a default value if not present
func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}

	return defaultValue
}

// LoadEnvVars will load a ".env[.development|.test]" file if it exists and set ENV vars.
// Useful in development and test modes. Not used in production.
func LoadEnvVariables() {

	env := GetEnv("GIN_ENV", "development")

	if env == "production" || env == "staging" {
		slog.Info("Not using .env file in production or staging.")
		return
	}

	filename := ".env." + env

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		filename = ".env"
	}

	err := godotenv.Load(filename)
	if err != nil {
		slog.Warn(".env file not loaded")
	}

}
