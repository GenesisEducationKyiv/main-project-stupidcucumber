package helpers

import (
	"api/bitcoin-api/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

type Rates struct {
}

type Rate struct {
}

const (
	httpsGeeko      = "https://api.coingecko.com"
	httpsGeekoRoute = "/api/v3/exchange_rates"
)

func RequestPriceGeeko() (float64, error) {
	u, err := url.ParseRequestURI(httpsGeeko)
	u.Path = httpsGeekoRoute
	finalUrl := fmt.Sprintf("%v", u)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error occured while parsing URI request: %v\n", err)
		return invalidPrice, err
	}

	exchangeRate, err := http.Get(finalUrl)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error occured while fetching price from Geeko: %v\n", err)
		return invalidPrice, err
	}

	defer exchangeRate.Body.Close()
	body, err := io.ReadAll(exchangeRate.Body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error occured while reading the request body: %v\n", err)
		return invalidPrice, err
	}

	var data models.Response

	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error occured while unmarshalling json: %v", err)
		return invalidPrice, err
	}

	// Extract the "value" field
	value := data.Rates["uah"].Value

	return value, nil
}
