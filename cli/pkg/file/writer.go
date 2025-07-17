package file

import (
	"os"
	"path/filepath"
)

// WriteFile writes data to the specified path creating directories as needed.
func WriteFile(path string, data []byte) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o600)
}
