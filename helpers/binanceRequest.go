package helpers

import (
	"api/bitcoin-api/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

func RequestPriceBinance() float64 {
	base := "https://api.binance.com"
	recource := "/api/v3/avgPrice"
	params := url.Values{}
	params.Add("symbol", "BTCUAH")

	u, _ := url.ParseRequestURI(base)
	u.Path = recource
	u.RawQuery = params.Encode()
	finalUrl := fmt.Sprintf("%v", u)

	exchangeRate, err := http.Get(finalUrl)

	if err != nil {
		fmt.Println(err.Error())
		return -1
	}

	defer exchangeRate.Body.Close()

	body, _ := ioutil.ReadAll(exchangeRate.Body)
	var exchangeRateObj models.ExchangeRate
	if err := json.Unmarshal(body, &exchangeRateObj); err != nil {
		fmt.Println(err.Error())
		return -1
	}

	result, _ := strconv.ParseFloat(exchangeRateObj.Price, 64)
	return result
}
