package integrationtests

import (
	"api/bitcoin-api/controller"
	"api/bitcoin-api/providers"
	"api/bitcoin-api/tests/mocks"
	"fmt"
	"testing"
	"time"
)

func TestGetPriceFirstFalse(t *testing.T) {
	mockCacheProvider := mocks.MockCacheProvider{
		Time: time.Now(),
	}
	mockPriceProviders := []providers.PriceProvider{
		&mocks.PriceProviderMock{Error: fmt.Errorf("service is down"), Price: -1},
		&mocks.PriceProviderMock{Error: nil, Price: 0.0},
	}

	price, err := controller.GetPrice(&mockCacheProvider, mockPriceProviders)

	if price != 0.0 {
		t.Errorf("expected price doesn't match actual: %f", price)
	}

	if err != nil {
		t.Errorf("unexpected error was thrown: %v", err)
	}
}

func TestGetPricAllFalse(t *testing.T) {

	mockCacheProvider := mocks.MockCacheProvider{
		Time: time.Time{},
	}
	mockPriceProviders := []providers.PriceProvider{
		&mocks.PriceProviderMock{Error: fmt.Errorf("service is down"), Price: -1},
		&mocks.PriceProviderMock{Error: fmt.Errorf("service is down"), Price: -1},
	}

	_, err := controller.GetPrice(&mockCacheProvider, mockPriceProviders)

	if err == nil {
		t.Errorf("err should not be nil")
	}
}
