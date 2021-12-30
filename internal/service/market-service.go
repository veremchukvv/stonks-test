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

// func (msi *MarketServiceImp) DeleteDeal(ctx context.Context, token string, dealID int) error {
//	return nil
// }
