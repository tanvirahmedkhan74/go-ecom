package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost string
	Port       string
	DBUser     string
	DBPassword string
	DBName     string
	SSLMode    string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()
	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "localhost"),
		Port:       getEnv("PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "12345"),
		DBName:     getEnv("DB_NAME", "ecom"),
		SSLMode:    getEnv("SSL_MODE", "disabled"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
