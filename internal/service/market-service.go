package service

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/veremchukvv/stonks-test/internal/models"
	"github.com/veremchukvv/stonks-test/internal/repository"
)

type MarketServiceImp struct {
	repo repository.MarketRepository
}

func NewMarketServiceImp(repo repository.MarketRepository) *MarketServiceImp {
	return &MarketServiceImp{
		repo,
	}
}

func (msi *MarketServiceImp) GetAllStocks(ctx context.Context) ([]*models.DealResp, error) {
	return msi.repo.GetAllStocks(ctx)
}

func (msi *MarketServiceImp) GetOneStock(ctx context.Context, stockID int) (*models.DealResp, error) {
	return msi.repo.GetOneStock(ctx, stockID)
}

func (msi *MarketServiceImp) CreateDeal(ctx context.Context, token string, stockID int, stockAmount int, portfolioID int) (int, error) {
	_, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(key *jwt.Token) (interface{}, error) {
		return []byte(SignKey), nil
	})
	if err != nil {
		return 0, err
	}

	return msi.repo.CreateDeal(ctx, stockID, stockAmount, portfolioID)
}

func (msi *MarketServiceImp) GetCurrencies(ctx context.Context) (*models.CurrencyRates, error) {
	resp, err := http.Get("https://iss.moex.com/iss/statistics/engines/currency/markets/selt/rates.json?iss.only=wap_rates")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	readResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var respJSON *models.Outer

	err = json.Unmarshal(readResp, &respJSON)
	if err != nil {
		log.Println("can't unmarshal response from IMOEX: ", err)
	}

	cr := &models.CurrencyRates{
		respJSON.Wap.Data[0][4].(float64),
		respJSON.Wap.Data[0][5].(float64),
		respJSON.Wap.Data[1][4].(float64),
		respJSON.Wap.Data[1][5].(float64),
	}
	return cr, nil
}

// func (msi *MarketServiceImp) DeleteDeal(ctx context.Context, token string, dealID int) error {
//	return nil
// }
