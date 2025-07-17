package util

import "testing"

func TestIsValidEmail(t *testing.T) {
	tests := []struct {
		email string
		valid bool
	}{
		{"john@example.com", true},
		{"john+label@domain.io", true},
		{"bad-email", false},
		{"@missing.com", false},
	}
	for _, tt := range tests {
		if got := IsValidEmail(tt.email); got != tt.valid {
			t.Errorf("IsValidEmail(%q) = %v, want %v", tt.email, got, tt.valid)
		}
	}
}
