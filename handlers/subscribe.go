package handlers

import (
	"api/bitcoin-api/controlers"
	"api/bitcoin-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostSubscribe(c *gin.Context) {
	var newEmail models.Email

	if err := c.BindJSON(&newEmail); err != nil {
		c.IndentedJSON(http.StatusConflict, newEmail)
		return
	}

	if err := controlers.AddEmail(newEmail); err == nil {
		c.IndentedJSON(http.StatusOK, newEmail)
	} else {
		c.IndentedJSON(http.StatusConflict, err.Error())
	}
}
