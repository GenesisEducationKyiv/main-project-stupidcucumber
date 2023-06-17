package handler

import (
	"net/http"

	"api/bitcoin-api/controller"

	"github.com/gin-gonic/gin"
)

func PostSendEmails(c *gin.Context) {
	price, _ := controller.GetPrice()

	controller.SendEmail(price)

	c.IndentedJSON(http.StatusOK, "Emails had been sent")
}
