package models

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type EmailCredentials struct {
	HostEmail    string
	HostPassword string
	HostSMTP     string
	PortSMTP     int
}

func NewEmailCredentials() (*EmailCredentials, error) {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	hostEmail := os.Getenv("HOST_EMAIL")
	hostPassword := os.Getenv("HOST_PASSWORD")
	hostSMTP := os.Getenv("SMTP_HOST")
	portSMTP := os.Getenv("SMTP_PORT")

	port, err := strconv.ParseInt(portSMTP, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("sending email: %w", err)
	}

	return &EmailCredentials{hostEmail, hostPassword, hostSMTP, int(port)}, nil
}
