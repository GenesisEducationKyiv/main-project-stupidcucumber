package helpers

import (
	"api/bitcoin-api/models"
	"net/mail"
)

func ValidateEmail(email models.Email) bool {
	if _, err := mail.ParseAddress(email.Email); err != nil {
		return false
	}

	return true
}
