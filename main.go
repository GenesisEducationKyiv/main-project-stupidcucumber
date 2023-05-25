package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getPrice(c *gin.Context) {
	var price float64 = 30

	c.IndentedJSON(200, price)
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

func main() {
	router := gin.Default()

	router.GET("/rate", getPrice)
	router.POST("/subscribe", postSubscribe)

	router.Run("localhost:8080")
}
