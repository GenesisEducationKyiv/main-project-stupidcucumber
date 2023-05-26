package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func requestPriceBinance() float64 {
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
	var exchangeRateObj ExchangeRate
	if err := json.Unmarshal(body, &exchangeRateObj); err != nil {
		fmt.Println(err.Error())
		return -1
	}

	result, _ := strconv.ParseFloat(exchangeRateObj.Price, 64)
	return result
}

func getPrice(c *gin.Context) {
	answer := make(map[string]float64)

	price := requestPriceBinance()
	answer["rate"] = price

	c.IndentedJSON(200, answer)
}

func postSubscribe(c *gin.Context) {
	var newEmail Email

	if err := c.BindJSON(&newEmail); err != nil {
		c.IndentedJSON(http.StatusConflict, newEmail)
		return
	}

	if !findEmail(newEmail) {
		c.IndentedJSON(409, newEmail)
		return
	}

	addEmail(newEmail)

	c.IndentedJSON(200, newEmail)
}

func postSendEmails(c *gin.Context) {
	price := requestPriceBinance()

	sendEmail(price)

	c.IndentedJSON(200, "Emails had been sent")
}

func main() {
	router := gin.Default()

	router.GET("/rate", getPrice)
	router.POST("/subscribe", postSubscribe)
	router.POST("/sendEmails", postSendEmails)

	router.Run("localhost:8080")
}
