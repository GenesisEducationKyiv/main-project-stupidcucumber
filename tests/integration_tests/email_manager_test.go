package integrationtests

import (
	"api/bitcoin-api/controller"
	"api/bitcoin-api/models"
	"api/bitcoin-api/tests/mocks"
	"fmt"
	"testing"
)

func TestSubscribeTrue(t *testing.T) {
	mockDatabase := mocks.DatabaseMock{
		Error:        nil,
		EmailPresent: false,
		Emails:       []models.Email{},
	}

	err := controller.Subscribe(models.Email{Email: "email2000@net.ua"}, &mockDatabase)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestSubscribeIsAlreadyPresent(t *testing.T) {
	present := models.Email{Email: "email2000@net.ua"}

	mockDatabase := mocks.DatabaseMock{
		Error:        nil,
		EmailPresent: true,
		Emails:       []models.Email{present},
	}

	err := controller.Subscribe(present, &mockDatabase)

	if err == nil {
		t.Errorf("expected error had not been returned! email %#v is present!", present)
	}
}

func TestSubscribeWithNoConnection(t *testing.T) {
	present := models.Email{Email: "email2000@net.ua"}

	mockDatabase := mocks.DatabaseMock{
		Error:        fmt.Errorf("database is down"),
		EmailPresent: false,
		Emails:       nil,
	}

	err := controller.Subscribe(present, &mockDatabase)
	if err == nil {
		t.Errorf("expected error, but nil is returned! actual database is down")
	}
}
