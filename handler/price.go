package handler

import (
	"fmt"
	"net/http"
	"os"

	"api/bitcoin-api/controller"
	"api/bitcoin-api/models"
	"api/bitcoin-api/providers"

	"github.com/gin-gonic/gin"
)

func GetPrice(c *gin.Context) {
	answer := make(map[string]float64)
	cacheProvider, err := models.NewFileCache()

	priceProviders := []providers.PriceProvider{&models.BinancePrice{}, &models.CoingeckoPrice{}}

	if err != nil {
		fmt.Fprintf(os.Stdout, "get price: %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
	}

	price, err := controller.GetPrice(cacheProvider, priceProviders)
	if err != nil {
		fmt.Fprintf(os.Stdout, "get price: %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	answer["rate"] = price

	c.IndentedJSON(http.StatusOK, answer)
}
