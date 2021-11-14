package models

type Portfolio struct {
	Id          int           `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Public      bool          `json:"public"`
	Stocks      map[Stock]int `json:"stocks"`
	Cash        float32       `json:"cash"`
}
