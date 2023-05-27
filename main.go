package main

import (
	"net/http"

	"log"

	"api/bitcoin-api/controlers"
	"api/bitcoin-api/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func getPrice(c *gin.Context) {
	answer := make(map[string]float64)

	price := controlers.GetPrice()
	answer["rate"] = price

	c.IndentedJSON(200, answer)
}

func postSubscribe(c *gin.Context) {
	var newEmail models.Email

	if err := c.BindJSON(&newEmail); err != nil {
		c.IndentedJSON(http.StatusConflict, newEmail)
		return
	}

	if !controlers.FindEmail(newEmail) {
		c.IndentedJSON(409, newEmail)
		return
	}

	err := controlers.AddEmail(newEmail)

	if err == nil {
		c.IndentedJSON(200, newEmail)
	} else {
		c.IndentedJSON(409, err.Error())
	}
}

func postSendEmails(c *gin.Context) {
	price := controlers.GetPrice()

	controlers.SendEmail(price)

	c.IndentedJSON(200, "Emails had been sent")
}

func main() {
	router := gin.Default()

	router.GET("/rate", getPrice)
	router.POST("/subscribe", postSubscribe)
	router.POST("/sendEmails", postSendEmails)

	router.Run("localhost:8080")
}
