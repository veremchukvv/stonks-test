package models

type Stock struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Summary  string  `json:"summary"`
	Type     string  `json:"type"`
	Cost     float32 `json:"cost"`
	Currency string  `json:"currency"`
}
