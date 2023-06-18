package controller

import (
	"api/bitcoin-api/interfaces"
	"api/bitcoin-api/models"
	"fmt"
	"os"
	"time"
)

func getPrice() (float64, error) {
	prices := []interfaces.Pricable{&models.BinancePrice{}, &models.CoingeckoPrice{}}

	for i := 0; i < len(prices); i++ {
		if price, err := prices[i].GetPrice(); err == nil {
			return price, nil
		}
	}

	return float64(invalidPrice), fmt.Errorf("Servers from where prices being fetched are down")
}

func updatePrice() (float64, error) {
	price, err := getPrice()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error occured while requesting price: %v\n", err)
		return float64(invalidPrice), err
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
		fmt.Fprintf(os.Stderr, "Error occured while reading cache: %v\n", err)
		return float64(invalidPrice), err
	}

	if time.Since(cache.TimeStamp).Minutes() <= 10 && time.Since(cache.TimeStamp).Hours() < 1 {
		return cache.Price, nil
	}

	new_price, err := updatePrice()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error occured while updating price: %v\n", err)
		return float64(invalidPrice), err
	}

	return new_price, nil
}
