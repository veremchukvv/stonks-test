package service

import (
	"context"
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

func (msi *MarketServiceImp) GetAllStocks(ctx context.Context) ([]*models.StockResp, error) {
	return msi.repo.GetAllStocks(ctx)
}

func (msi *MarketServiceImp) GetOneStock(ctx context.Context, stockId int) (*models.StockResp, error) {
	return msi.repo.GetOneStock(ctx, stockId)
}

func (msi *MarketServiceImp) CreateDeal(ctx context.Context, token string, stockId int, stockAmount int, portfolioId int) (int, error) {

	_, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(key *jwt.Token) (interface{}, error) {
		return []byte(SignKey), nil
	})
	if err != nil {
		return 0, err
	}

	return msi.repo.CreateDeal(ctx, stockId, stockAmount, portfolioId)
}

func (msi *MarketServiceImp) DeleteDeal(ctx context.Context, token string, dealId int) error {
	return nil
}
