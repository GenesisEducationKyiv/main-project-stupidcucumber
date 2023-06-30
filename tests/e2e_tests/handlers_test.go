package e2etests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

func TestSubscribeHandler(t *testing.T) {
	expectedStatusCode := http.StatusOK
	jsonData := []byte(`{"email":"test@email.com"}`)

	response, err := http.Post("http://localhost:8080/api/subscribe", "application/json", bytes.NewReader(jsonData))
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	defer response.Body.Close()

	if response.StatusCode != expectedStatusCode {
		t.Errorf("status code is different from what is expected: %v", response.StatusCode)
	}
}

func TestPriceHandler(t *testing.T) {
	response, err := http.Get("http://localhost:8080/api/rate")
	if err != nil {
		t.Errorf("unexpected error occured: %v", err)
	}

	var price struct {
		Rate float64 `json:"rate"`
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Errorf("unexpected error during reading from response: %v", err)
	}

	err = json.Unmarshal(body, &price)
	if err != nil {
		t.Errorf("unexpected error during unmarshalling body of the response: %v", err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		t.Errorf("status code is different from what is expected!")
	}

	if price.Rate < 0 {
		t.Errorf("unexpected value")
	}
}
