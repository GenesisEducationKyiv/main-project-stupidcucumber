package models

type Response struct {
	Rates map[string]struct {
		Name  string  `json:"name"`
		Unit  string  `json:"unit"`
		Value float64 `json:"value"`
		Type  string  `json:"type"`
	} `json:"rates"`
}
