package config

// config.go loads environment variables and exposes configuration structs.
// Configurations should be reusable and provided via interfaces.

type AppConfig struct {
	// TODO: define configuration fields, e.g., server port
}

// Load reads environment variables into Config structs.
func Load() (*AppConfig, error) {
	// TODO: load and return configuration
	return &AppConfig{}, nil
}
