package unittest

import (
	"api/bitcoin-api/models"
	"testing"
)

func TestValidateEmail(t *testing.T) {
	emails := []models.Email{
		{Email: "surname2000@mail.com"},
		{Email: "abv"},
	}
	expected := []bool{true, false}

	for index, email := range emails {
		result := email.Validate()

		if result != expected[index] {
			t.Errorf("got %t, but expected %t, on %v", result, expected, email)
		}
	}
}
