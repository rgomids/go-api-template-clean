package config

import "testing"

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

func TestLoadConfigFull(t *testing.T) {
	t.Setenv("APP_ENV", "prod")
	t.Setenv("PORT", "9090")
	t.Setenv("DATABASE_URL", "db")
	t.Setenv("SMTP_PORT", "2525")
	cfg, err := LoadConfig()
	if err != nil {
		t.Fatalf("load error: %v", err)
	}
	if !cfg.IsProd() || cfg.Port != "9090" || cfg.SMTPPort != 2525 {
		t.Fatal("values not loaded correctly")
	}
}
