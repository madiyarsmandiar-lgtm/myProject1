package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port       string
	DBHost     string
	DBUser     string
	DBPassword string
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		Port:       GetEnv("PORT", "8080"),
		DBHost:     GetEnv("DB_HOST", "localhost"),
		DBUser:     GetEnv("DB_USER", "Madi"),
		DBPassword: GetEnv("DB_PASSWORD", "Almaty"),
	}, nil
}

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
