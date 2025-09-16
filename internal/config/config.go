package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT string
}

func Load() (*Config, error) {
	err := godotenv.Load(".env")

	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	config := &Config{
		PORT: getEnv("PORT"),
	}

	return config, nil
}

func getEnv(key string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	panic("required environment variable " + key + " is not set")
}
