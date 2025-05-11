package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv read file .env
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found, using system environtment variables")
	}
}

// GetEnv take value form environment variable
func GetEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
