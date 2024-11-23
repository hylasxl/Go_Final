package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	config := &Config{
		DBUser:     getEnv("DB_USER", "flight_management_user"),
		DBPassword: getEnv("DB_PASSWORD", "LOgbBjq58lYaKccwLYZCYWqAUmwPxYOh"),
		DBName:     getEnv("DB_NAME", "flight_management"),
		DBHost:     getEnv("DB_HOST", "dpg-csu6fea3esus73c82i60-a.singapore-postgres.render.com"),
		DBPort:     getEnv("DB_PORT", "5432"),
	}

	return config
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
