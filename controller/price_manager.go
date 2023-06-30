package controller

import (
	"api/bitcoin-api/models"
	"api/bitcoin-api/providers"
	"fmt"
	"time"
)

var (
	invalidPrice = -1
)

func getPrice(priceProviders []providers.PriceProvider) (float64, error) {
	for i := 0; i < len(priceProviders); i++ {
		if price, err := priceProviders[i].GetPrice(); err == nil {
			return price, nil
		}
	}

	return float64(invalidPrice), fmt.Errorf("servers from where prices being fetched are down")
}

func updatePrice(cacheProvider providers.CacheProvider, priceProviders []providers.PriceProvider) (float64, error) {
	price, err := getPrice(priceProviders)
	if err != nil {
		return float64(invalidPrice), fmt.Errorf("updating price: %w", err)
	}

	cache := models.CachedPrice{
		TimeStamp: time.Now(),
		Price:     price,
	}

	err = cacheProvider.Write(cache)
	if err != nil {
		return 0.0, fmt.Errorf("updating price: %w", err)
	}

	return price, nil
}

func GetPrice(cacheProvider providers.CacheProvider, priceProviders []providers.PriceProvider) (float64, error) {
	cache, err := cacheProvider.Read()

	if err != nil {
		return float64(invalidPrice), fmt.Errorf("getting price: %w", err)
	}

	if time.Since(cache.TimeStamp).Minutes() <= 10 && time.Since(cache.TimeStamp).Hours() < 1 {
		return cache.Price, nil
	}

	newPrice, err := updatePrice(cacheProvider, priceProviders)
	if err != nil {
		return float64(invalidPrice), fmt.Errorf("getting price: %w", err)
	}

	return newPrice, nil
}
