package models

//type Data struct {
//	TradeDate       string
//	TradeTime       string
//	SecId           string
//	ShortName       string
//	Price           float32
//	LastToPrevPrice float32
//	Nominal         int
//	Decimals        int
//}
//
//type Datas struct {
//	Datas []Data `json:"data"`
//}

type Inner struct {
	Data [][]interface{} `json:"data"`
}

type Outer struct {
	Wap Inner `json:"wap_rates"`
}

//func (d *Datas) UnmarshalJSON(buf []byte) error {
//	// tmp := []interface{}{&d.TradeDate, &d.TradeTime, &d.SecId, &d.ShortName, &d.Price, &d.LastToPrevPrice, &d.Nominal, &d.Decimals}
//	fmt.Println("hello!")
//	// tmp := []interface{}{&d.Datas[].TradeDate, &d.Datas[1].TradeTime, &d.Datas[1].SecId, &d.Datas[1].ShortName, &d.Datas[1].Price, &d.Datas[1].LastToPrevPrice, &d.Datas[1].Nominal, &d.Datas[1].Decimals}
//	tmp := []interface{}{&d.Datas}
//	wantLen := len(tmp)
//
//	if err := json.Unmarshal(buf, &tmp); err != nil {
//		return err
//	}
//	if g, e := len(tmp), wantLen; g != e {
//		return fmt.Errorf("wrong number of fields in Notification: %d != %d", g, e)
//	}
//	return nil
//}

type CurrencyRates struct {
	USDRate   float64 `json:"usd_rate"`
	USDChange float64 `json:"usd_change"`
	EURRate   float64 `json:"eur_rate"`
	EURChange float64 `json:"eur_change"`
}
