package main

import (
	"fmt"
	"os"
	"strings"
)

func getEmails() []string {
	file, err := os.ReadFile("email.db")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while extracting emails: %v", err)
		return nil
	}

	return strings.Split(string(file), "\n")
}

func findEmail(email Email) bool {
	var emails []string = getEmails()

	for i := 0; i < len(emails); i++ {
		if emails[i] == email.Email {
			return false
		}
	}

	return true
}

func addEmail(email Email) {
	f, err := os.OpenFile("email.db", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0744)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer f.Close()

	f.WriteString(email.Email + "\n")
}
