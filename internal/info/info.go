package info

import (
	"os"
	"strings"
)

// FilePath specifies the path of the VERSION file. It can be overridden in tests.
var FilePath = "VERSION"

// ReadVersion returns the application version read from FilePath.
func ReadVersion() (string, error) {
	data, err := os.ReadFile(FilePath)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(data)), nil
}
