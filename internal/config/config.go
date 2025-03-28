package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config holds application configuration
type Config struct {
	Port           string
	Environment    string
	AppName        string
	DBHost         string
	DBPort         string
	DBUser         string
	DBPassword     string
	DBName         string
	DBSSLMode      string
	DATABASE_URL   string
	EmailHost      string
	EmailPort      string
	EmailUsername  string
	EmailPassword  string
	EmailFrom      string
	JWTSecret      string
	PasswordSalt   string
	LogLevel       string
	AllowedOrigins string
	GinMode        string
}

// Load loads configuration from environment variables
func Load() (*Config, error) {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	cfg := &Config{
		Port:           getEnv("PORT", ":8080"),
		Environment:    getEnv("ENVIRONMENT", "development"),
		AppName:        getEnv("APP_NAME", "goUniAdmin"),
		DBSSLMode:      getEnv("DB_SSLMODE", "disable"),
		DATABASE_URL:   getEnv("DATABASE_URL", ""),
		EmailHost:      getEnv("EMAIL_HOST", "smtp.example.com"),
		EmailPort:      getEnv("EMAIL_PORT", "587"),
		EmailUsername:  getEnv("EMAIL_USERNAME", "noreply@example.com"),
		EmailPassword:  getEnv("EMAIL_PASSWORD", "emailsecret"),
		EmailFrom:      getEnv("EMAIL_FROM", "noreply@example.com"),
		JWTSecret:      getEnv("JWT_SECRET", "your-very-secret-key-here"),
		PasswordSalt:   getEnv("PASSWORD_SALT", "some-random-salt"),
		LogLevel:       getEnv("LOG_LEVEL", "debug"),
		AllowedOrigins: getEnv("ALLOWED_ORIGINS", "http://localhost:3000"),
		GinMode:        getEnv("GIN_MODE", "debug"),
	}

	return cfg, nil
}

// getEnv retrieves an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
