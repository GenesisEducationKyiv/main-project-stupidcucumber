package models

import (
	"api/bitcoin-api/tools/loaders"
	"fmt"
	"os"
	"strings"
)

type FileDatabase struct {
	FileName string `json:"filename"`
}

func NewFileDatabase() (*FileDatabase, error) {
	name, err := loaders.GetEnvVariable("DATABASE_PATH")
	if err != nil {
		return nil, fmt.Errorf("instantiating new DataBase Connector: %w", err)
	}

	if _, err := os.Stat(name); err != nil {
		_, err := os.Create(name)
		if err != nil {
			return nil, fmt.Errorf("creating new database: %w", err)
		}
	}

	return &FileDatabase{name}, nil
}

func (db *FileDatabase) ReadAll() ([]Email, error) {
	file, err := os.ReadFile(db.FileName)
	if err != nil {
		return nil, fmt.Errorf("reading file %s: %w", db.FileName, err)
	}

	unfilteredEmails := strings.Split(string(file), "\n")
	filteredEmails := []Email{}

	for i := 0; i < len(unfilteredEmails); i++ {
		if unfilteredEmails[i] != "" {
			filteredEmails = append(filteredEmails, Email{unfilteredEmails[i]})
		}
	}

	return filteredEmails, nil
}

func (db *FileDatabase) IsPresent(email Email) (bool, error) {
	emails, err := db.ReadAll()
	if err != nil {
		return false, fmt.Errorf("checking whether email is present: %w", err)
	}

	for _, dbEmail := range emails {
		if dbEmail == email {
			return true, nil
		}
	}

	return false, nil
}

func (db *FileDatabase) Write(email Email) error {
	f, err := os.OpenFile(db.FileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return fmt.Errorf("opening the file %s: %w", db.FileName, err)
	}

	defer f.Close()

	isPresent, err := db.IsPresent(email)
	if err != nil {
		return fmt.Errorf("adding email %s: %w", email.Email, err)
	}

	if !email.Validate() || isPresent {
		return fmt.Errorf("provided email is invalid")
	}

	if _, err := f.WriteString(email.Email + "\n"); err != nil {
		return fmt.Errorf("the writing to the file went wrong: %w", err)
	}

	return nil
}
