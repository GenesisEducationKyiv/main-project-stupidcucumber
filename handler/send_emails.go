package handler

import (
	"fmt"
	"net/http"
	"os"

	"api/bitcoin-api/controller"

	"github.com/gin-gonic/gin"
)

func PostSendEmails(c *gin.Context) {
	price, _ := controller.GetPrice()

	err := controller.SendEmail(price)
	if err != nil {
		fmt.Fprintf(os.Stderr, "sending emails: %v", err)
	}

	c.IndentedJSON(http.StatusOK, "Emails had been sent")
}
