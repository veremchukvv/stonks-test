package models

type Stock struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Type        string  `json:"type"`
	Cost        float32 `json:"cost"`
	Currency    string  `json:"currency"`
}
