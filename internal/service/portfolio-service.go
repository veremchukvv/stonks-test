package service

import (
	"context"
	"github.com/veremchukvv/stonks-test/internal/models"
	"github.com/veremchukvv/stonks-test/internal/repository"
)

type PortfolioServiceImp struct {
	repo repository.PortfolioRepository
}

func NewPortfolioServiceImp(repo repository.PortfolioRepository) *PortfolioServiceImp {
	return &PortfolioServiceImp{repo}
}

func (ps *PortfolioServiceImp) GetAllPortfolios(ctx context.Context) ([]*models.Portfolio, error) {
	return nil, nil
}

func (ps *PortfolioServiceImp) GetOnePortfolio(ctx context.Context, portfolio_id int) (*models.Portfolio, error) {
	return nil, nil
}

func (ps *PortfolioServiceImp) CreatePortfolio(ctx context.Context, portfolio *models.Portfolio) (*models.Portfolio, error) {
	return nil, nil
}

func (ps *PortfolioServiceImp) DeletePortfolio(ctx context.Context, portfolio_id int) error {
	return nil
}
