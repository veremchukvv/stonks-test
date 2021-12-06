package repository

import (
	"context"
	"github.com/veremchukvv/stonks-test/internal/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) (*models.User, error)
	DeleteUser(ctx context.Context, userId int, authType string) error
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	CreateVKUser(ctx context.Context, user *models.User) (*models.User, error)
	GetVKUserByID(ctx context.Context, vkid int) (*models.User, error)
	GetUserByID(ctx context.Context, userId int, authType string) (*models.User, error)
}

type PortfolioRepository interface {
	GetAllPortfolios(ctx context.Context, userId int, authType string) ([]*models.Portfolio, error)
	GetPortfolioDeals(ctx context.Context, portfolioId int) (*models.OnePortfolioResp, []*models.DealResp, error)
	GetPortfolioClosedDeals(ctx context.Context, portfolioId int) ([]*models.DealResp, error)
	CreatePortfolio(ctx context.Context, userId int, authType string, portfolio *models.Portfolio) (*models.Portfolio, error)
	DeletePortfolio(ctx context.Context, portfolioId int) error
}

type MarketRepository interface {
	GetAllStocks(ctx context.Context) ([]*models.DealResp, error)
	GetOneStock(ctx context.Context, stockId int) (*models.DealResp, error)
	CreateDeal(ctx context.Context, stockId int, stockAmount int, portfolioId int) (int, error)
}

type DealRepository interface {
	GetOneDeal(ctx context.Context, dealId int) (*models.DealResp, error)
	CloseDeal(ctx context.Context, dealId int) error
	DeleteDeal(ctx context.Context, dealId int) error
	GetOneClosedDeal(ctx context.Context, closedDealId int) (*models.DealResp, error)
	DeleteClosedDeal(ctx context.Context, closedDealId int) error
}