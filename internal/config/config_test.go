package config

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadConfigDefaults(t *testing.T) {
	t.Setenv("DATABASE_URL", "db")
	cfg, err := LoadConfig()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.Env != "dev" {
		t.Errorf("expected dev env, got %s", cfg.Env)
	}
	if cfg.Port != "8080" {
		t.Errorf("expected default port 8080, got %s", cfg.Port)
	}
	if cfg.SMTPPort != 587 {
		t.Errorf("expected default smtp port 587, got %d", cfg.SMTPPort)
	}
}

func TestLoadConfigMissingDB(t *testing.T) {
	if _, err := LoadConfig(); err == nil {
		t.Fatal("expected error for missing DATABASE_URL")
	}
}

func TestLoadConfigInvalidSMTPPort(t *testing.T) {
	t.Setenv("DATABASE_URL", "db")
	t.Setenv("SMTP_PORT", "bad")
	if _, err := LoadConfig(); err == nil {
		t.Fatal("expected error for invalid SMTP_PORT")
	}
}

func TestLoadEnvFile(t *testing.T) {
	dir := t.TempDir()
	envPath := filepath.Join(dir, ".env")
	os.WriteFile(envPath, []byte("SOME_VAR=value\n"), 0o600)
	t.Setenv("SOME_VAR", "")

	if err := loadEnvFile(envPath); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if v := os.Getenv("SOME_VAR"); v != "value" {
		t.Errorf("expected SOME_VAR to be 'value', got '%s'", v)
	}
}

func TestLoadEnvFileMissing(t *testing.T) {
	if err := loadEnvFile("no-file.env"); err != nil {
		t.Fatalf("expected no error for missing file, got %v", err)
	}
}

func TestLoadConfigLoadsDotEnv(t *testing.T) {
	dir := t.TempDir()
	envPath := filepath.Join(dir, ".env")
	os.WriteFile(envPath, []byte("DATABASE_URL=db\nAPP_ENV=prod\nPORT=9090\nSMTP_PORT=2525\n"), 0o600)

	t.Setenv("DATABASE_URL", "")
	t.Setenv("APP_ENV", "")
	t.Setenv("PORT", "")
	t.Setenv("SMTP_PORT", "")

	cwd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(cwd)

	cfg, err := LoadConfig()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.DatabaseURL != "db" {
		t.Errorf("expected db from .env, got %s", cfg.DatabaseURL)
	}
	if cfg.Env != "prod" {
		t.Errorf("expected prod env, got %s", cfg.Env)
	}
	if cfg.Port != "9090" {
		t.Errorf("expected port 9090, got %s", cfg.Port)
	}
	if cfg.SMTPPort != 2525 {
		t.Errorf("expected smtp port 2525, got %d", cfg.SMTPPort)
	}
}

func TestIsProd(t *testing.T) {
	cfg := &AppConfig{Env: "prod"}
	if !cfg.IsProd() {
		t.Errorf("expected IsProd to return true")
	}
	cfg.Env = "dev"
	if cfg.IsProd() {
		t.Errorf("expected IsProd to return false")
	}
}

func TestLoadEnvFileStatError(t *testing.T) {
	path := strings.Repeat("a", 5000)
	if err := loadEnvFile(path); err == nil {
		t.Fatal("expected error for stat failure")
	}
}

func TestLoadEnvFileParseError(t *testing.T) {
	dir := t.TempDir()
	envPath := filepath.Join(dir, "bad.env")
	os.WriteFile(envPath, []byte("INVALID LINE\nfoo=bar"), 0o600)
	if err := loadEnvFile(envPath); err == nil {
		t.Fatal("expected error for invalid env file")
	}
}
