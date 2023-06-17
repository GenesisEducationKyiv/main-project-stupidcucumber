package handlers

import (
	"net/http"

	"api/bitcoin-api/controlers"

	"github.com/gin-gonic/gin"
)

func GetPrice(c *gin.Context) {
	answer := make(map[string]float64)

	price, err := controlers.GetPrice()
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	answer["rate"] = price

	c.IndentedJSON(http.StatusOK, answer)
}
