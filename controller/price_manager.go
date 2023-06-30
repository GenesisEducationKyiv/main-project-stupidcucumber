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

func getPrice() (float64, error) {
	prices := []providers.PriceProvider{&models.BinancePrice{}, &models.CoingeckoPrice{}}

	for i := 0; i < len(prices); i++ {
		if price, err := prices[i].GetPrice(); err == nil {
			return price, nil
		}
	}

	return float64(invalidPrice), fmt.Errorf("servers from where prices being fetched are down")
}

func updatePrice() (float64, error) {
	price, err := getPrice()
	if err != nil {
		return float64(invalidPrice), fmt.Errorf("updating price: %w", err)
	}

	cache := models.CachedPrice{
		TimeStamp: time.Now(),
		Price:     price,
	}

	writeCache(cache)

	return price, nil
}

func GetPrice() (float64, error) {
	cache, err := readCache()
	if err != nil {
		return float64(invalidPrice), fmt.Errorf("getting price: %w", err)
	}

	if time.Since(cache.TimeStamp).Minutes() <= 10 && time.Since(cache.TimeStamp).Hours() < 1 {
		return cache.Price, nil
	}

	newPrice, err := updatePrice()
	if err != nil {
		return float64(invalidPrice), fmt.Errorf("getting price: %w", err)
	}

	return newPrice, nil
}
