package controlers

import (
	"fmt"
	"os"
	"strings"

	"api/bitcoin-api/models"
)

func getEmails() []string {
	file, err := os.ReadFile("email.db")

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
	f, err := os.OpenFile("email.db", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0744)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer f.Close()

	f.WriteString(email.Email + "\n")
}
