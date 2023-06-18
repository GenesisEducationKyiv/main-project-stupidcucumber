package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

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

func writeCache(cache models.CachedPrice) {
	cachedJSON, err := json.Marshal(cache)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error occured during jsonifying CachedPrice: %v\n", err)
		return
	}

	if err = os.WriteFile(CACHE_PATH, cachedJSON, 0o766); err != nil {
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
