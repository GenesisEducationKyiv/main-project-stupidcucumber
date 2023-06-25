package handler

import (
	"fmt"
	"net/http"
	"os"

	"api/bitcoin-api/controller"
	"api/bitcoin-api/models"

	"github.com/gin-gonic/gin"
)

func PostSendEmails(c *gin.Context) {
	cacheProvider, err := models.NewFileCache()
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	database, err := models.NewFileDatabase()
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	emailCredentials, err := models.NewEmailCredentials()
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	price, err := controller.GetPrice(cacheProvider)
	if err != nil {
		fmt.Fprintf(os.Stderr, "sending emails: %v", err)
		return
	}

	err = controller.SendEmail(price, database, emailCredentials)
	if err != nil {
		fmt.Fprintf(os.Stderr, "sending emails: %v", err)
	}

	c.IndentedJSON(http.StatusOK, "Emails had been sent")
}
