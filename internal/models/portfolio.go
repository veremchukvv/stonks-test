package models

type Portfolio struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Public      bool    `json:"is_public"`
	AssetsRUB   float64 `json:"assets_rub"   ticker:"RUB"`
	AssetsUSD   float64 `json:"assets_usd"   ticker:"USD"`
	AssetsEUR   float64 `json:"assets_eur"   ticker:"EUR"`
	ProfitRUB   float64 `json:"profit_rub"`
	ProfitUSD   float64 `json:"profit_usd"`
	ProfitEUR   float64 `json:"profit_eur"`
	PercentRUB  float64 `json:"percent_rub"`
	PercentUSD  float64 `json:"percent_usd"`
	PercentEUR  float64 `json:"percent_eur"`
	// Stocks      map[Stock]int        `json:"stocks"`
	// Cash        map[Currency]float32 `json:"cash"`
}

type OnePortfolioResp struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Public      bool   `json:"is_public"`
}
