package controlers

import (
	"fmt"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func generateMessage(to string, price float64) *gomail.Message {
	message := gomail.NewMessage()
	message.SetHeader("From", os.Getenv("HOST_EMAIL"))
	message.SetHeader("To", to)
	message.SetHeader("Subject", "Cryptocurrency rate to UAH")
	message.SetBody("text/plain", fmt.Sprintf("Current price of BTC in Binance is %f", price))

	return message
}

func SendEmail(price float64) {
	port, _ := strconv.ParseInt(os.Getenv("SMTP_PORT"), 10, 64)
	dialer := gomail.NewDialer(os.Getenv("SMTP_HOST"), int(port),
		os.Getenv("HOST_EMAIL"), os.Getenv("HOST_PASSWORD"))

	fmt.Println("Trying to log in!")

	emails := getEmails()

	for i := 0; i < len(emails); i++ {
		var message gomail.Message = *generateMessage(emails[i], price)

		if err := dialer.DialAndSend(&message); err != nil {
			fmt.Fprintf(os.Stderr, "Error while sending to %s occured: %v", emails[i], err)
		}
	}
}
