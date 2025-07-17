package file

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteFileSuccess(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "test.txt")
	if err := WriteFile(path, []byte("data")); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	b, err := os.ReadFile(path)
	if err != nil || string(b) != "data" {
		t.Fatalf("file not written correctly")
	}
}

func TestWriteFileBadPath(t *testing.T) {
	if err := WriteFile("", []byte("data")); err == nil {
		t.Fatal("expected error for bad path")
	}
}
