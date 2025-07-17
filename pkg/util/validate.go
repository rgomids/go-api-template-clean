package util

import "regexp"

// IsValidEmail performs a minimal regex validation for email addresses.
func IsValidEmail(email string) bool {
	pattern := `^[a-zA-Z0-9_.%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	return regexp.MustCompile(pattern).MatchString(email)
}
