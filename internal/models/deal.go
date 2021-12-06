package models

import "time"

type DealReq struct {
	StockID int `json:"stock_id"`
	StockAmount int `json:"stock_amount"`
	PortfolioID int `json:"portfolio_id"`
}

type DealResp struct {
	Id             int       `json:"id"`
	Ticker         string    `json:"ticker"`
	Name           string    `json:"name"`
	Type           string    `json:"type"`
	Description    string    `json:"description"`
	Amount         int       `json:"amount"`
	Cost           float32   `json:"cost"`
	Value          float32   `json:"value"`
	Currency       string    `json:"currency"`
	OpenedAt       time.Time `json:"opened_at"`
	ClosedAt       time.Time `json:"closed_at"`
	Profit         float32   `json:"profit"`
	Percent        float32   `json:"percent"`

}