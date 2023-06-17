package handlers

import (
	"net/http"

	"api/bitcoin-api/controlers"

	"github.com/gin-gonic/gin"
)

func PostSendEmails(c *gin.Context) {
	price, _ := controlers.GetPrice()

	controlers.SendEmail(price)

	c.IndentedJSON(http.StatusOK, "Emails had been sent")
}
