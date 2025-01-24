package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PostgresHost     string
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string
	PostgresPort     string
	ServerPort       string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found, relying on environment variables")
	}

	return &Config{
		PostgresHost:     getEnv("POSTGRES_HOST", "localhost"),
		PostgresUser:     getEnv("POSTGRES_USER", "user"),
		PostgresPassword: getEnv("POSTGRES_PASSWORD", "password"),
		PostgresDB:       getEnv("POSTGRES_DB", "chatroom"),
		PostgresPort:     getEnv("POSTGRES_PORT", "5422"),
		ServerPort:       getEnv("SERVER_PORT", "8080"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
