package models

import "encoding/json"

type Currency struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Ticker   string  `json:"ticker"`
	BuyRate  float32 `json:"buy_rate"`
	SellRate float32 `json:"sell_rate"`
}

func (s Currency) MarshalText() (text []byte, err error) {
	type x Currency
	return json.Marshal(x(s))
}
