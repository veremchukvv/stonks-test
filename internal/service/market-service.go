package service

import (
	"context"
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

func (msi *MarketServiceImp) GetAllStocks(ctx context.Context) ([]*models.Stock, error) {
	return msi.repo.GetAllStocks(ctx)
}

func (msi *MarketServiceImp) GetOneStock(ctx context.Context, stockId int) (*models.Stock, error) {
	return nil, nil
}

func (msi *MarketServiceImp) CreateDeal(ctx context.Context, token string, stockId int, stockAmount int) (int, error) {
	return 0, nil
}

func (msi *MarketServiceImp) DeleteDeal(ctx context.Context, token string, dealId int) error {
	return nil
}
