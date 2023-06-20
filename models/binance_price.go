package models

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type BinancePrice struct {
	Price string `json:"price"`
}

const (
	httpsBinance       = "https://api.binancee.com"
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
	finalURL := fmt.Sprintf("%v", u)

	if err != nil {
		return invalidPrice, fmt.Errorf("Error occured while parsing the URL: %w", err)
	}

	exchangeRate, err := http.Get(finalURL)
	if err != nil {
		return invalidPrice, fmt.Errorf("Error occured while requesting GET from the %s: %w",
			httpsBinance+httpsBinanceRoute, err)
	}

	defer exchangeRate.Body.Close()

	body, err := io.ReadAll(exchangeRate.Body)
	if err != nil {
		return invalidPrice, fmt.Errorf("Error occured while reading the request body: %w", err)
	}

	if err := json.Unmarshal(body, p); err != nil {
		return invalidPrice, fmt.Errorf("unmarshalling exchange rate object: %w", err)
	}

	result, _ := strconv.ParseFloat(p.Price, 64)
	return result, nil
}
