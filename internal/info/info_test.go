package info

import (
	"os"
	"path/filepath"
	"testing"
)

func TestReadVersionSuccess(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "VERSION")
	os.WriteFile(path, []byte("1.2.3\n"), 0o600)

	old := FilePath
	FilePath = path
	defer func() { FilePath = old }()

	v, err := ReadVersion()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if v != "1.2.3" {
		t.Errorf("expected version '1.2.3', got '%s'", v)
	}
}

func TestReadVersionMissingFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "missing")

	old := FilePath
	FilePath = path
	defer func() { FilePath = old }()

	if _, err := ReadVersion(); err == nil {
		t.Fatal("expected error for missing file")
	}
}
