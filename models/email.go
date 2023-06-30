package models

import "net/mail"

type Email struct {
	Email string `json:"email"`
}

func (e *Email) Validate() bool {
	if _, err := mail.ParseAddress(e.Email); err != nil {
		return false
	}

	return true
}
