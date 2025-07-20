package util

import "github.com/google/uuid"

// GenerateID returns a new UUID string.
func GenerateID() string {
	return uuid.NewString()
}
