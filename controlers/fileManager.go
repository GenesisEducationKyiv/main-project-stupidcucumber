package controlers

import (
	"fmt"
	"log"
	"os"
	"strings"

	"api/bitcoin-api/helpers"
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

func findEmail(email models.Email) bool {
	var emails []string = getEmails()

	for i := 0; i < len(emails); i++ {
		if emails[i] == email.Email {
			return true
		}
	}

	return false
}

func AddEmail(email models.Email) error {
	f, err := os.OpenFile(DATABASE_PATH,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0744)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error occured while reading adding the Email: %v", err)
		return err
	}

	fmt.Printf("Email added: %s", email.Email)

	defer f.Close()

	if helpers.ValidateEmail(email) && !findEmail(email) {
		f.WriteString(email.Email + "\n")
	} else {
		return fmt.Errorf("provided email is invalid")
	}

	return nil
}
