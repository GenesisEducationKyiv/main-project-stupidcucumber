package controller

import (
	"fmt"
	"os"
	"strings"

	"api/bitcoin-api/helpers"
	"api/bitcoin-api/models"
)

func getEmails() ([]string, error) {
	path, err := helpers.GetEnvVariable("DATABASE_PATH")
	if err != nil {
		return nil, fmt.Errorf("extracting emails: %w", err)
	}

	file, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading file %s: %w", path, err)
	}

	unfilteredEmails := strings.Split(string(file), "\n")
	filteredEmails := []string{}

	for i := 0; i < len(unfilteredEmails); i++ {
		if unfilteredEmails[i] != "" {
			filteredEmails = append(filteredEmails, unfilteredEmails[i])
		}
	}

	return filteredEmails, nil
}

func findEmail(email models.Email) (bool, error) {
	emails, err := getEmails()
	if err != nil {
		return false, fmt.Errorf("searching for the email %s: %w", email.Email, err)
	}

	for i := 0; i < len(emails); i++ {
		if emails[i] == email.Email {
			return true, nil
		}
	}

	return false, nil
}

func AddEmail(email models.Email) error {
	path, err := helpers.GetEnvVariable("DATABASE_PATH")
	if err != nil {
		return fmt.Errorf("getting .env variable: %w", err)
	}

	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return fmt.Errorf("opening the file %s: %w", path, err)
	}

	defer f.Close()

	isPresent, err := findEmail(email)
	if err != nil {
		return fmt.Errorf("adding email %s: %w", email.Email, err)
	}

	if !helpers.ValidateEmail(email.Email) || isPresent {
		return fmt.Errorf("provided email is invalid")
	}

	if _, err := f.WriteString(email.Email + "\n"); err != nil {
		return fmt.Errorf("the writing to the file went wrong: %w", err)
	}

	return nil
}
