package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds application configuration
type Config struct {
	Port                 string
	Environment          string
	AppName              string
	DBHost               string
	DBPort               string
	DBUser               string
	DBPassword           string
	DBName               string
	DBSSLMode            string
	DATABASE_URL         string
	EmailHost            string
	EmailPort            string
	EmailUsername        string
	EmailPassword        string
	EmailFrom            string
	JWTSecret            string
	PasswordSalt         string
	LogLevel             string
	AllowedOrigins       string
	GinMode              string
	SwaggerHost          string // Swagger-specific host
	IsHTTPAuthForSwagger bool   // Enable HTTP basic auth for Swagger
	SwaggerAuthUser      string // Username for Swagger auth
	SwaggerAuthPassword  string // Password for Swagger auth
}

// ConfigInstance is a global instance of the configuration
var ConfigInstance *Config

// Load loads configuration from environment variables
func Load() (*Config, error) {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	cfg := &Config{
		Port:                 getEnv("PORT", ":8080"),
		Environment:          getEnv("ENVIRONMENT", "development"),
		AppName:              getEnv("APP_NAME", "goUniAdmin"),
		DBSSLMode:            getEnv("DB_SSLMODE", "disable"),
		DATABASE_URL:         getEnv("DATABASE_URL", ""),
		EmailHost:            getEnv("EMAIL_HOST", "smtp.example.com"),
		EmailPort:            getEnv("EMAIL_PORT", "587"),
		EmailUsername:        getEnv("EMAIL_USERNAME", "noreply@example.com"),
		EmailPassword:        getEnv("EMAIL_PASSWORD", "emailsecret"),
		EmailFrom:            getEnv("EMAIL_FROM", "noreply@example.com"),
		JWTSecret:            getEnv("JWT_SECRET", "your-very-secret-key-here"),
		PasswordSalt:         getEnv("PASSWORD_SALT", "some-random-salt"),
		LogLevel:             getEnv("LOG_LEVEL", "debug"),
		AllowedOrigins:       getEnv("ALLOWED_ORIGINS", "http://localhost:3000"),
		GinMode:              getEnv("GIN_MODE", "debug"),
		SwaggerHost:          getEnv("SWAGGER_HOST", "localhost:8080"),
		IsHTTPAuthForSwagger: getEnvAsBool("IS_HTTP_AUTH_FOR_SWAGGER", true),
		SwaggerAuthUser:      getEnv("SWAGGER_AUTH_USER", "indianic"),
		SwaggerAuthPassword:  getEnv("SWAGGER_AUTH_PASSWORD", "indianic"),
	}

	ConfigInstance = cfg
	return cfg, nil
}

// getEnv retrieves an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// getEnvAsBool retrieves an environment variable as a boolean or returns a default value
func getEnvAsBool(key string, defaultValue bool) bool {
	if value, exists := os.LookupEnv(key); exists {
		b, err := strconv.ParseBool(value)
		if err == nil {
			return b
		}
	}
	return defaultValue
}

// GetSwaggerHost returns the Swagger host (for compatibility with previous method)
func (c *Config) GetSwaggerHost() string {
	return c.SwaggerHost
}
