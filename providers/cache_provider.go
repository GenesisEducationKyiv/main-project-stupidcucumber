package providers

import (
	"api/bitcoin-api/models"
)

type CacheProvider interface {
	Read() (*models.CachedPrice, error)
	Write(models.CachedPrice) error
}
