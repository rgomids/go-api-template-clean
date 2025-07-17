package version

import (
	"os"
	"strings"
)

// Load reads the version from the given file.
func Load(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(data)), nil
}
