package controlers

import (
	"fmt"
	"log"
	"os"
	"strings"

	"api/bitcoin-api/models"

	"github.com/joho/godotenv"
)

var (
	DATABASE_PATH = os.Getenv("DATABASE_PATH")
)

func init() {

	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	DATABASE_PATH = os.Getenv("DATABASE_PATH")
}

func getEmails() []string {
	file, err := os.ReadFile(DATABASE_PATH)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while extracting emails: %v", err)
		return nil
	}

	unfilteredEmails := strings.Split(string(file), "\n")
	filteredEmails := []string{}

	for i := 0; i < len(unfilteredEmails); i++ {
		if unfilteredEmails[i] != "" {
			filteredEmails = append(filteredEmails, unfilteredEmails[i])
		}
	}

	return filteredEmails
}

func FindEmail(email models.Email) bool {
	var emails []string = getEmails()

	for i := 0; i < len(emails); i++ {
		if emails[i] == email.Email {
			return false
		}
	}

	return true
}

func AddEmail(email models.Email) {
	f, err := os.OpenFile(DATABASE_PATH,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0744)

	//TODO: add pattern recognition to validate registration of incoming emails

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error occured while reading adding the Email: %v", err)
		return
	}

	defer f.Close()

	f.WriteString(email.Email + "\n")
}
