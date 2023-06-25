package handler

import (
	"fmt"
	"net/http"
	"os"

	"api/bitcoin-api/controller"

	"github.com/gin-gonic/gin"
)

func PostSendEmails(c *gin.Context) {
	cacheProvider, err := controller.NewFileCache()
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	price, err := controller.GetPrice(cacheProvider)
	if err != nil {
		fmt.Fprintf(os.Stderr, "sending emails: %v", err)
		return
	}

	err = controller.SendEmail(price)
	if err != nil {
		fmt.Fprintf(os.Stderr, "sending emails: %v", err)
	}

	c.IndentedJSON(http.StatusOK, "Emails had been sent")
}
