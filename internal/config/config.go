package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT        string
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	SSL_MODE    string
	JWT_SECRET  string
}

func Load() (*Config, error) {
	err := godotenv.Load(".env")

	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	config := &Config{
		PORT:        getEnv("PORT"),
		DB_HOST:     getEnv("DB_HOST"),
		DB_PORT:     getEnv("DB_PORT"),
		DB_USER:     getEnv("DB_USER"),
		DB_PASSWORD: getEnv("DB_PASSWORD"),
		DB_NAME:     getEnv("DB_NAME"),
		SSL_MODE:    getEnv("SSL_MODE"),
		JWT_SECRET:  getEnv("JWT_SECRET"),
	}

	return config, nil
}

func getEnv(key string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	panic("required environment variable " + key + " is not set")
}
