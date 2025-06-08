package config

import (
	"os"

	_ "github.com/joho/godotenv"
)

type Config struct {
	AppEnv       string
	AppPort      string
	AppURL       string
	DBHost       string
	DBPort       string
	DBUser       string
	DBPassword   string
	DBName       string
	DBSSLMode    string
	JWTSecret    string
	JWTExpiresIn string
}

func LoadConfig() *Config {
	return &Config{
		AppEnv:       getEnv("APP_ENV", "development"),
		AppPort:      getEnv("APP_PORT", "3000"),
		AppURL:       getEnv("APP_URL", "http://localhost:3000"),
		DBHost:       getEnv("DB_HOST", "localhost"),
		DBPort:       getEnv("DB_PORT", "5432"),
		DBUser:       getEnv("DB_USER", "postgres"),
		DBPassword:   getEnv("DB_PASS", "123456"),
		DBName:       getEnv("DB_NAME", "fiberecomapidb"),
		DBSSLMode:    getEnv("DB_SSL", "disable"),
		JWTSecret:    getEnv("JWT_SECRET", "fibernextcommerce_jwt_secret_key_2024"),
		JWTExpiresIn: getEnv("JWT_EXPIRES_IN", "24h"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}