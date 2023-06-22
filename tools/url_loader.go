package tools

import (
	"fmt"
	"io"
	"net/http"
)

func LoadURL(url string) ([]byte, error) {
	result, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error occured while requesting GET from the %s: %w", url, err)
	}

	defer result.Body.Close()

	body, err := io.ReadAll(result.Body)
	if err != nil {
		return nil, fmt.Errorf("error occured while reading the request body: %w", err)
	}

	return body, nil
}
