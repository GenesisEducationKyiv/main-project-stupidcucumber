package models

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

type CoingeckoPrice struct {
	Rates map[string]struct {
		Name  string  `json:"name"`
		Unit  string  `json:"unit"`
		Value float64 `json:"value"`
		Type  string  `json:"type"`
	} `json:"rates"`
}

const (
	httpsGeeko      = "https://api.coingecko.com"
	httpsGeekoRoute = "/api/v3/exchange_rates"
)

func (p *CoingeckoPrice) GetPrice() (float64, error) {
	u, err := url.ParseRequestURI(httpsGeeko)
	u.Path = httpsGeekoRoute
	finalURL := fmt.Sprintf("%s", u)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error occured while parsing URI request: %v\n", err)
		return invalidPrice, err
	}

	exchangeRate, err := http.Get(finalURL)
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

	err = json.Unmarshal(body, p)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error occured while unmarshalling json: %v", err)
		return invalidPrice, err
	}

	// Extract the "value" field
	value := p.Rates["uah"].Value

	return value, nil
}
