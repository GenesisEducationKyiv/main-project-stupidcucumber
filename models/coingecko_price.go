package models

import (
	"api/bitcoin-api/helpers"
	"encoding/json"
	"fmt"
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
	body, err := helpers.LoadURL(httpsGeeko + httpsGeekoRoute)
	if err != nil {
		return invalidPrice, fmt.Errorf("getting coingecko price: %w", err)
	}

	err = json.Unmarshal(body, p)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error occured while unmarshalling json: %v", err)
		return invalidPrice, err
	}

	value := p.Rates["uah"].Value

	return value, nil
}
