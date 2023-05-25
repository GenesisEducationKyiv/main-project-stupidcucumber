package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func findEmail(email Email) bool {

	f, err := os.ReadFile("email.db")

	if err != nil && errors.Is(err, os.ErrNotExist) {
		fmt.Println(err.Error())
		return true
	} else if err != nil {
		fmt.Println(err.Error())
		return false
	}

	var emails []string = strings.Split(string(f), "\n")

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
