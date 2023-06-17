package helpers

import (
	"net/mail"

	"api/bitcoin-api/models"
)

func ValidateEmail(email models.Email) bool {
	if _, err := mail.ParseAddress(email.Email); err != nil {
		return false
	}

	return true
}
