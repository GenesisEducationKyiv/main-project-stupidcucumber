package validators

import "testing"

func TestValidateEmail(t *testing.T) {
	emails := []string{"surname2000@mail.com", "abv"}
	expected := []bool{true, false}

	for index, email := range emails {
		result := ValidateEmail(email)

		if result != expected[index] {
			t.Errorf("got %t, but expected %t, on %s", result, expected, email)
		}
	}
}
