package models

import (
	"encoding/json"
)

type Stock struct {
	Id          int     `json:"id"`
	Ticker      string  `json:"ticker"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Type        string  `json:"type"`
	Cost        float32 `json:"cost"`
	Currency    string  `json:"currency"`
}

func (s Stock) MarshalText() (text []byte, err error) {
	type x Stock
	return json.Marshal(x(s))
}
