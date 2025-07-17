package entity_test

import (
	"testing"

	"github.com/rgomids/go-api-template-clean/internal/domain/entity"
)

func TestUserIsValidEmail(t *testing.T) {
	cases := []struct {
		email string
		valid bool
	}{
		{"test@example.com", true},
		{"user+label@domain.co", true},
		{"invalid-email", false},
		{"@missinguser.com", false},
	}

	for _, c := range cases {
		u := &entity.User{Email: c.email}
		if got := u.IsValidEmail(); got != c.valid {
			t.Errorf("IsValidEmail(%q) = %v, want %v", c.email, got, c.valid)
		}
	}
}
