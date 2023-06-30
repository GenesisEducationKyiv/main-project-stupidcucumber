package controller

import (
	"api/bitcoin-api/models"
	"api/bitcoin-api/providers"
	"bytes"
	"fmt"
	"os"
	"text/template"

	"gopkg.in/gomail.v2"
)

func generateMessage(to string, price float64, emailCredentials *models.EmailCredentials) (*gomail.Message, error) {
	t, err := template.ParseFiles("templates/template.html")
	if err != nil {
		return nil, fmt.Errorf("generating message: %w", err)
	}
	var body bytes.Buffer

	if err := t.Execute(&body, struct {
		Rate string
	}{
		Rate: fmt.Sprintf("%f", price),
	}); err != nil {
		return nil, fmt.Errorf("error occured during generating message: %w", err)
	}

	message := gomail.NewMessage()
	message.SetHeader("From", emailCredentials.HostEmail)
	message.SetHeader("To", to)
	message.SetHeader("Subject", "Cryptocurrency rate to UAH")
	message.Embed("templates/logo.png")
	message.Embed("templates/icons8-bitcoin-250.png")
	message.SetBody("text/html", body.String())

	return message, nil
}

func SendEmail(price float64, database providers.RepositoryProvider, emailCredentials *models.EmailCredentials) error {
	emails, err := database.ReadAll()
	if err != nil {
		return fmt.Errorf("sending emails: %w", err)
	}

	dialer := gomail.NewDialer(emailCredentials.HostSMTP, emailCredentials.PortSMTP,
		emailCredentials.HostEmail, emailCredentials.HostPassword)

	for i := 0; i < len(emails); i++ {
		message, err := generateMessage(emails[i].Email, price, emailCredentials)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error while sending to %s occured: %v", emails[i], err)
		}

		if err := dialer.DialAndSend(message); err != nil {
			fmt.Fprintf(os.Stderr, "error while sending to %s occured: %v", emails[i], err)
		}
	}

	return nil
}
