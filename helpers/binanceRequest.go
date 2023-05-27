package helpers

import (
	"api/bitcoin-api/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

const (
	httpsBinance       = "https://api.binance.com"
	httpsRoute         = "/api/v3/avgPrice"
	convertionCurrency = "BTCUAH"
)

func RequestPriceBinance() float64 {
	params := url.Values{}
	params.Add("symbol", convertionCurrency)

	u, _ := url.ParseRequestURI(httpsBinance)
	u.Path = httpsRoute
	u.RawQuery = params.Encode()
	finalUrl := fmt.Sprintf("%v", u)

	exchangeRate, err := http.Get(finalUrl)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error occured while requesting GET from the %s: %v",
			httpsBinance+httpsRoute, err)
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
