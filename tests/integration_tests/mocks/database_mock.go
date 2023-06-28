package mocks

import (
	"api/bitcoin-api/models"
	"fmt"
)

type DatabaseMock struct {
	EmailPresent bool
	Error        error
	Emails       []models.Email
}

func (d *DatabaseMock) IsPresent(models.Email) (bool, error) {
	return d.EmailPresent, d.Error
}

func (d *DatabaseMock) Write(models.Email) error {
	if d.EmailPresent {
		return fmt.Errorf("email already present")
	}

	if d.Error != nil {
		return d.Error
	}

	return nil
}

func (d *DatabaseMock) ReadAll() ([]models.Email, error) {
	return d.Emails, d.Error
}
