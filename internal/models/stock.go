package models

import "encoding/json"

type Stock struct {
	Id          int     `json:"id"`
	Ticker      string  `json:"ticker"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Type        string  `json:"type"`
	Cost        float32 `json:"cost"`
	Currency    string  `json:"currency"`
}

type StockResp struct {
	Id       int     `json:"id"`
	Ticker   string  `json:"ticker"`
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	Amount   int     `json:"amount"`
	Cost     float32 `json:"cost"`
	Value    float32 `json:"value"`
	Currency string  `json:"currency"`
}

func (s Stock) MarshalText() (text []byte, err error) {
	type x Stock
	return json.Marshal(x(s))
}
