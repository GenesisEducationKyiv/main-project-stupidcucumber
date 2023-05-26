package controlers

import (
	"bytes"
	"fmt"
	"os"
	"strconv"

	"text/template"

	"gopkg.in/gomail.v2"
)

func generateMessage(to string, price float64) *gomail.Message {
	t, _ := template.ParseFiles("templates/template.html")
	var body bytes.Buffer

	t.Execute(&body, struct {
		Rate string
	}{
		Rate: fmt.Sprintf("%f", price),
	})

	message := gomail.NewMessage()
	message.SetHeader("From", os.Getenv("HOST_EMAIL"))
	message.SetHeader("To", to)
	message.SetHeader("Subject", "Cryptocurrency rate to UAH")
	message.Embed("templates/logo.png")
	message.Embed("templates/icons8-bitcoin-250.png")
	message.SetBody("text/html", body.String())

	return message
}

func SendEmail(price float64) {
	port, _ := strconv.ParseInt(os.Getenv("SMTP_PORT"), 10, 64)
	dialer := gomail.NewDialer(os.Getenv("SMTP_HOST"), int(port),
		os.Getenv("HOST_EMAIL"), os.Getenv("HOST_PASSWORD"))

	emails := getEmails()

	for i := 0; i < len(emails); i++ {
		var message gomail.Message = *generateMessage(emails[i], price)

		if err := dialer.DialAndSend(&message); err != nil {
			fmt.Fprintf(os.Stderr, "Error while sending to %s occured: %v", emails[i], err)
		}
	}
}
