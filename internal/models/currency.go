package models

type Currency struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	BuyRate  float32 `json:"buy_rate"`
	SellRate float32 `json:"sell_rate"`
}
