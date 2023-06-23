package models

import (
	"api/bitcoin-api/tools/loaders"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

type BinancePrice struct {
	Price string `json:"price"`
}

const (
	httpsBinance       = "https://api.binance.com"
	httpsBinanceRoute  = "/api/v3/avgPrice"
	convertionCurrency = "BTCUAH"
	invalidPrice       = -1
)

func (p *BinancePrice) GetPrice() (float64, error) {
	params := url.Values{}
	params.Add("symbol", convertionCurrency)

	u, err := url.ParseRequestURI(httpsBinance)
	u.Path = httpsBinanceRoute
	u.RawQuery = params.Encode()
	finalURL := u.String()

	if err != nil {
		return invalidPrice, fmt.Errorf("error occured while parsing the URL: %w", err)
	}

	body, err := loaders.LoadURL(finalURL)
	if err != nil {
		return invalidPrice, fmt.Errorf("requesting URL: %w", err)
	}

	if err := json.Unmarshal(body, p); err != nil {
		return invalidPrice, fmt.Errorf("unmarshalling exchange rate object: %w", err)
	}

	result, _ := strconv.ParseFloat(p.Price, 64)
	return result, nil
}
