package handler

import (
	"net/http"

	"api/bitcoin-api/controller"
	"api/bitcoin-api/models"

	"github.com/gin-gonic/gin"
)

func PostSubscribe(c *gin.Context) {
	database, err := models.NewFileDatabase()
	if err != nil {
		c.IndentedJSON(http.StatusConflict, err.Error())
		return
	}

	var newEmail models.Email

	if err := c.BindJSON(&newEmail); err != nil {
		c.IndentedJSON(http.StatusConflict, newEmail)
		return
	}

	if err := controller.Subscribe(newEmail, database); err == nil {
		c.IndentedJSON(http.StatusOK, newEmail)
	} else {
		c.IndentedJSON(http.StatusConflict, err.Error())
	}
}
