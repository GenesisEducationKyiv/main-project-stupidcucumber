package validators

import (
	"net/mail"
)

func ValidateEmail(email string) bool {
	if _, err := mail.ParseAddress(email); err != nil {
		return false
	}

	return true
}
