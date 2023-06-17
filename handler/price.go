package handler

import (
	"net/http"

	"api/bitcoin-api/controller"

	"github.com/gin-gonic/gin"
)

func GetPrice(c *gin.Context) {
	answer := make(map[string]float64)

	price, err := controller.GetPrice()
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	answer["rate"] = price

	c.IndentedJSON(http.StatusOK, answer)
}
