package controlers

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"api/bitcoin-api/helpers"
	"api/bitcoin-api/models"
)

func init() {
	file, err := os.Create(".cache")
	file.Close()

	updatePrice()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error occured during creation of the file: %v", err)
	}
}

func writeCache(cache models.CachedPrice) {
	cached_json, err := json.Marshal(cache)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error occured during jsonifying CachedPrice: %v", err)
		return
	}

	err = os.WriteFile(".cache", cached_json, 0766)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error occured during writing to a file '.cache': %v", err)
		return
	}
}

func readCache() *models.CachedPrice {
	fileContent, err := os.ReadFile(".cache")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error occured during reading from '.cache': %v", err)
		return nil
	}

	var cache models.CachedPrice

	err = json.Unmarshal(fileContent, &cache)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error occured during unmarshalling the cache: %v", err)
		return nil
	}

	return &cache
}

func updatePrice() float64 {
	price := helpers.RequestPriceBinance()

	cache := models.CachedPrice{
		TimeStamp: time.Now(),
		Price:     price,
	}

	writeCache(cache)

	return price
}

func GetPrice() float64 {
	cache := readCache()

	fmt.Print(time.Now().Sub(cache.TimeStamp).Minutes())
	fmt.Print(time.Now().Sub(cache.TimeStamp).Hours())

	if time.Now().Sub(cache.TimeStamp).Minutes() <= 10 && time.Now().Sub(cache.TimeStamp).Hours() < 1 {
		fmt.Println("Wow this priced was cached!")
		return cache.Price
	} else {
		new_price := updatePrice()

		return new_price
	}
}
