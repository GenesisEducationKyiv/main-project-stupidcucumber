package models

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EmailCredentials struct {
	HostEmail    string
	HostPassword string
	HostSMTP     string
	PortSMTP     string
}

func NewEmailCredentials() *EmailCredentials {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	hostEmail := os.Getenv("HOST_EMAIL")
	hostPassword := os.Getenv("HOST_PASSWORD")
	hostSMTP := os.Getenv("SMTP_HOST")
	portSMTP := os.Getenv("SMTP_PORT")

	return &EmailCredentials{hostEmail, hostPassword, hostSMTP, portSMTP}
}
