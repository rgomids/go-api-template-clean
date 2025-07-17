package config

import "os"

// config.go loads environment variables and exposes configuration structs.
// Configurations should be reusable and provided via interfaces.

// AppConfig holds basic application configuration such as the HTTP server
// address. Additional configuration fields can be added as the project grows.
type AppConfig struct {
	ServerAddr string
}

// Load reads environment variables and returns an AppConfig populated with
// sensible defaults when values are not provided.
func Load() (*AppConfig, error) {
	addr := os.Getenv("HTTP_ADDR")
	if addr == "" {
		addr = ":8080"
	}
	return &AppConfig{ServerAddr: addr}, nil
}
