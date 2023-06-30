package controller

import (
	"fmt"

	"api/bitcoin-api/models"
	"api/bitcoin-api/providers"
)

func Subscribe(email models.Email, database providers.RepositoryProvider) error {
	isPresent, err := database.IsPresent(email)
	if err != nil {
		return fmt.Errorf("subscribing an email %s: %w", email.Email, err)
	}

	if isPresent {
		return fmt.Errorf("email is already present")
	}

	if err = database.Write(email); err != nil {
		return fmt.Errorf("writing an email to the database: %w", err)
	}

	return nil
}
