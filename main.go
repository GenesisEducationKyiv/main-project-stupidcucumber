package main

import (
	"os"

	"log"

	"api/bitcoin-api/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	router := gin.Default()

	router.GET("/rate", handlers.GetPrice)
	router.POST("/subscribe", handlers.PostSubscribe)
	router.POST("/sendEmails", handlers.PostSendEmails)

	router.Run(":" + os.Getenv("APP_PORT"))
}
