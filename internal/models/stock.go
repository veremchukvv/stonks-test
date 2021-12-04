package models

import (
	"encoding/json"
	"time"
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

type StockResp struct {
	Id             int       `json:"id"`
	Ticker         string    `json:"ticker"`
	Name           string    `json:"name"`
	Type           string    `json:"type"`
	Description    string    `json:"description"`
	Amount         int       `json:"amount"`
	Cost           float32   `json:"cost"`
	Value          float32   `json:"value"`
	Currency       string    `json:"currency"`
	CreatedAt      time.Time `json:"created_at"`
	IsClosed       bool      `json:"is_closed"`
	ClosingTime    time.Time `json:"closing_time"`
	ProfitCurrent  float32   `json:"profit_current"`
	PercentCurrent float32   `json:"percent_current"`
	ProfitClosed   float32   `json:"profit_closed"`
	PercentClosed  float32   `json:"percent_closed"`
}

func (s Stock) MarshalText() (text []byte, err error) {
	type x Stock
	return json.Marshal(x(s))
}
