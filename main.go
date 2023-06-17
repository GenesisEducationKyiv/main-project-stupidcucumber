package main

import (
	"log"
	"os"

	"api/bitcoin-api/handler"

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

	router.GET("/api/rate", handler.GetPrice)
	router.POST("/api/subscribe", handler.PostSubscribe)
	router.POST("/api/sendEmails", handler.PostSendEmails)

	log.Fatal(router.Run(":" + os.Getenv("APP_PORT")))
}
