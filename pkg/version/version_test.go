package version

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadSuccess(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "VERSION")
	os.WriteFile(f, []byte("1.2.3\n"), 0o600)
	v, err := Load(f)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if v != "1.2.3" {
		t.Errorf("expected version 1.2.3, got %s", v)
	}
}

func TestLoadError(t *testing.T) {
	if _, err := Load("nonexistent"); err == nil {
		t.Fatal("expected error for missing file")
	}
}
