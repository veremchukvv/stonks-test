package models

type Deal struct {
	StockID int `json:"stock_id"`
	StockAmount int `json:"stock_amount"`
	PortfolioID int `json:"portfolio_id"`
}
