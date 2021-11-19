package models

type Portfolio struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Public      bool    `json:"is_public"`
	AssetsRUB   float32 `json:"assets_rub"`
	AssetsUSD   float32 `json:"assets_usd"`
	AssetsEUR   float32 `json:"assets_eur"`
	//Stocks      map[Stock]int        `json:"stocks"`
	//Cash        map[Currency]float32 `json:"cash"`
}
