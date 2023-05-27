package controlers

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
	CACHE_PATH = os.Getenv("CACHE_PATH")
)

func init() {

	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	CACHE_PATH = os.Getenv("CACHE_PATH")
	file, err := os.Create(CACHE_PATH)
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

	err = os.WriteFile(CACHE_PATH, cached_json, 0766)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error occured during writing to a file '.cache': %v", err)
		return
	}
}

func readCache() *models.CachedPrice {

	fmt.Println(CACHE_PATH)
	fileContent, err := os.ReadFile(CACHE_PATH)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error occured during reading from %s: %v", CACHE_PATH, err)
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

	if time.Now().Sub(cache.TimeStamp).Minutes() <= 10 && time.Now().Sub(cache.TimeStamp).Hours() < 1 {
		return cache.Price
	} else {
		new_price := updatePrice()

		return new_price
	}
}
