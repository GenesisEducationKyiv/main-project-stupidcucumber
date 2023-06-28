package mocks

import (
	"api/bitcoin-api/models"
	"time"
)

type MockCacheProvider struct {
	Time time.Time
}

func (mock *MockCacheProvider) Read() (*models.CachedPrice, error) {
	return &models.CachedPrice{TimeStamp: mock.Time, Price: 0.0}, nil
}

func (mock *MockCacheProvider) Write(toWrite models.CachedPrice) error {
	return nil
}
