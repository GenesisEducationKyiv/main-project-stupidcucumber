package controller

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"api/bitcoin-api/models"
	"api/bitcoin-api/tools"
)

func writeCache(cache models.CachedPrice) error {
	path, err := tools.GetEnvVariable("CACHE_PATH")
	if err != nil {
		return fmt.Errorf("getting .env vavriable in writeCahce: %w", err)
	}

	cachedJSON, err := json.Marshal(cache)
	if err != nil {
		return fmt.Errorf("jsonifying CachedPrice: %w", err)
	}

	if err = os.WriteFile(path, cachedJSON, 0o766); err != nil {
		return fmt.Errorf("writing to the file %s: %w", path, err)
	}

	return nil
}

func readCache() (*models.CachedPrice, error) {
	path, err := tools.GetEnvVariable("CACHE_PATH")
	if err != nil {
		return nil, fmt.Errorf("reading cache: %w", err)
	}

	fileContent, err := os.ReadFile(path)

	if err != nil && os.IsNotExist(err) {
		price, err := getPrice()
		if err != nil {
			return &models.CachedPrice{}, err
		}
		file, _ := os.Create(path)
		defer file.Close()

		writeCache(models.CachedPrice{Price: price, TimeStamp: time.Now()})

		return &models.CachedPrice{Price: price, TimeStamp: time.Now()}, nil
	}

	var cache models.CachedPrice

	if err = json.Unmarshal(fileContent, &cache); err != nil {
		return nil, fmt.Errorf("unmarshalling JSON object into CachedPrice: %w", err)
	}

	return &cache, nil
}
