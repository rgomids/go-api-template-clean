package config

import (
	"fmt"
	"os"
	"strconv"
)

// AppConfig holds all configurable parameters of the application. It is loaded
// from environment variables in a centralized manner so the rest of the
// application can simply depend on this struct.
type AppConfig struct {
	Env          string
	Port         string
	DatabaseURL  string
	RedisURL     string
	SMTPHost     string
	SMTPPort     int
	SMTPUser     string
	SMTPPassword string
}

// LoadConfig reads the required environment variables, sets default values for
// optional ones and validates mandatory fields. It returns an AppConfig instance
// ready to be consumed by the application.
func LoadConfig() (*AppConfig, error) {
	cfg := &AppConfig{}

	// Environment the application is running in. Defaults to "dev" when not
	// provided.
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}
	cfg.Env = env

	// HTTP port the server should listen on. Defaults to 8080.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	cfg.Port = port

	// Mandatory fields.
	cfg.DatabaseURL = os.Getenv("DATABASE_URL")
	if cfg.DatabaseURL == "" {
		return nil, fmt.Errorf("DATABASE_URL is required")
	}

	cfg.RedisURL = os.Getenv("REDIS_URL")

	cfg.SMTPHost = os.Getenv("SMTP_HOST")

	if v := os.Getenv("SMTP_PORT"); v != "" {
		p, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("invalid SMTP_PORT: %w", err)
		}
		cfg.SMTPPort = p
	} else {
		cfg.SMTPPort = 587
	}

	cfg.SMTPUser = os.Getenv("SMTP_USER")
	cfg.SMTPPassword = os.Getenv("SMTP_PASSWORD")

	return cfg, nil
}

// IsProd indicates if the application is running in production environment.
func (c *AppConfig) IsProd() bool {
	return c.Env == "prod"
}
