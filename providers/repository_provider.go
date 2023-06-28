package providers

import "api/bitcoin-api/models"

type RepositoryProvider interface {
	Write(models.Email) error
	ReadAll() ([]models.Email, error)
	IsPresent(models.Email) (bool, error)
}
