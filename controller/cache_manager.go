package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"api/bitcoin-api/helpers"
	"api/bitcoin-api/models"

	"github.com/joho/godotenv"
)

var (
	CACHE_PATH   = os.Getenv("CACHE_PATH")
	invalidPrice = -1
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	CACHE_PATH = os.Getenv("CACHE_PATH")
}

func getPrice() (float64, error) {
	price, err := helpers.RequestPriceBinance()

	if err == nil {
		return price, nil
	}

	price, err = helpers.RequestPriceGeeko()

	return price, err
}

func writeCache(cache models.CachedPrice) {
	cached_json, err := json.Marshal(cache)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error occured during jsonifying CachedPrice: %v\n", err)
		return
	}

	if err = os.WriteFile(CACHE_PATH, cached_json, 0o766); err != nil {
		fmt.Fprintf(os.Stderr, "Error occured during writing to a file '.cache': %v\n", err)
		return
	}
}

func readCache() (*models.CachedPrice, error) {
	fileContent, err := os.ReadFile(CACHE_PATH)

	if err != nil && os.IsNotExist(err) {
		price, err := getPrice()
		if err != nil {
			return &models.CachedPrice{}, err
		}
		CACHE_PATH = os.Getenv("CACHE_PATH")
		file, _ := os.Create(CACHE_PATH)
		file.Close()

		writeCache(models.CachedPrice{Price: price, TimeStamp: time.Now()})

		return &models.CachedPrice{Price: price, TimeStamp: time.Now()}, nil
	}

	var cache models.CachedPrice

	if err = json.Unmarshal(fileContent, &cache); err != nil {
		fmt.Fprintf(os.Stderr, "Error occured during unmarshalling the cache: %v\n", err)
		return nil, err
	}

	return &cache, nil
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
