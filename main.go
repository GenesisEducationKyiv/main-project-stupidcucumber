package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func postSubscribe(c *gin.Context) {
	var newEmail Email

	if err := c.BindJSON(&newEmail); err != nil {
		return
	}

	c.IndentedJSON(http.StatusCreated, newEmail)
}

func main() {
	router := gin.Default()

	router.POST("/subscribe", postSubscribe)

	router.Run("localhost:8080")
}
