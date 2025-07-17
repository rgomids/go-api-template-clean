package entity

import (
	"regexp"
	"time"
)

// User represents a system user with basic attributes.
type User struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
}

// IsValidEmail performs a basic regex validation of the user's email.
func (u *User) IsValidEmail() bool {
	pattern := `^[a-zA-Z0-9_.%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	return regexp.MustCompile(pattern).MatchString(u.Email)
}
