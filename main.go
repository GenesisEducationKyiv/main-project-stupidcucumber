package main

import (
	"log"
	"os"

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

	router.GET("/api/rate", handlers.GetPrice)
	router.POST("/api/subscribe", handlers.PostSubscribe)
	router.POST("/api/sendEmails", handlers.PostSendEmails)

	log.Fatal(router.Run(":" + os.Getenv("APP_PORT")))
}
