package unittest

import (
	"api/bitcoin-api/models"
	"bytes"
	"fmt"
	"os"
	"testing"
)

var (
	expectedTestEmail = models.Email{
		Email: "new.email1000@mail.testnet",
	}

	testFileNameDatabase = "testEmail.db"
)

func cleanupTestsFileDatabase() {
	if _, err := os.Stat(testFileNameDatabase); err == nil {
		if err := os.Remove(testFileNameDatabase); err != nil {
			fmt.Fprintf(os.Stderr, "removing file: %v", err)
		}
	}
}

func TestFileDatabaseWrite(t *testing.T) {
	fileDatabase := models.FileDatabase{
		FileName: testFileNameDatabase,
	}

	if err := fileDatabase.Write(expectedTestEmail); err != nil {
		t.Errorf("unexpected error during writing to the test database file: %v", err)
	}

	fileBody, err := os.ReadFile(testFileNameDatabase)
	if err != nil {
		t.Errorf("unexpected error during reading the test database file: %v", err)
	}

	actualTestEmail := string(bytes.SplitAfter(fileBody, []byte("\n"))[0])
	if expectedTestEmail.Email == actualTestEmail {
		t.Errorf("expected email (%s) is different from actual (%s)", expectedTestEmail.Email, actualTestEmail)
	}

	t.Cleanup(cleanupTestsFileDatabase)
}

func TestFileDatabaseReadAll(t *testing.T) {
	fileDatabase := models.FileDatabase{
		FileName: testFileNameDatabase,
	}

	expectedTestEmails := []models.Email{
		{Email: "mymail@john.com"},
		{Email: "hellomyfriend@ukr.net"},
		expectedTestEmail,
	}

	for _, email := range expectedTestEmails {
		if err := fileDatabase.Write(email); err != nil {
			t.Errorf("unexpected error while writing %#v to the test database file: %v", email, err)
		}
	}

	fileBody, err := os.ReadFile(testFileNameDatabase)
	if err != nil {
		t.Errorf("unexpected error while reading the test database file: %v", err)
	}

	actualTestEmails := bytes.Split(fileBody, []byte("\n"))
	actualTestEmails = actualTestEmails[:len(actualTestEmails)-1]

	for index, byteEmail := range actualTestEmails {
		email := models.Email{
			Email: string(byteEmail),
		}
		if expectedTestEmails[index] != email {
			t.Errorf("actual email is different from expected")
		}
	}

	t.Cleanup(cleanupTestsFileDatabase)
}

func TestFileDatabaseIsPresent(t *testing.T) {
	fileDatabase := models.FileDatabase{
		FileName: testFileNameDatabase,
	}

	expectedTestEmails := []models.Email{
		{Email: "mymail@john.com"},
		{Email: "hellomyfriend@ukr.net"},
		expectedTestEmail,
	}

	for _, email := range expectedTestEmails {
		if err := fileDatabase.Write(email); err != nil {
			t.Errorf("unexpected error while writing %#v to the test database file: %v", email, err)
		}
	}

	expectedTestAnswer := true

	actualTestAnswer, err := fileDatabase.IsPresent(expectedTestEmail)
	if err != nil {
		t.Errorf("unexpected error during checking whether the email is present: %v", err)
	}

	if actualTestAnswer != expectedTestAnswer {
		t.Errorf("actual answer is different from the expected")
	}

	t.Cleanup(cleanupTestsFileDatabase)
}
