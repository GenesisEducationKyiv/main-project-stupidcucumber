package handlers

import (
	"api/bitcoin-api/controlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostSendEmails(c *gin.Context) {
	price, _ := controlers.GetPrice()

	controlers.SendEmail(price)

	c.IndentedJSON(http.StatusOK, "Emails had been sent")
}
