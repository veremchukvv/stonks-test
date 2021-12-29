package repository

import (
	"context"

	"github.com/veremchukvv/stonks-test/internal/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) (*models.User, error)
	DeleteUser(ctx context.Context, id int, authType string) error
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	CreateVKUser(ctx context.Context, user *models.User) (*models.User, error)
	GetVKUserByID(ctx context.Context, id int) (*models.User, error)
	CreateGoogleUser(ctx context.Context, user *models.User) (*models.User, error)
	GetGoogleUserByID(ctx context.Context, gid string) (*models.User, error)
	GetUserByID(ctx context.Context, id int, authType string) (*models.User, error)
}

type PortfolioRepository interface {
	GetAllPortfolios(ctx context.Context, userId int, authType string) ([]*models.Portfolio, error)
	GetPortfolioDeals(ctx context.Context, portfolioID int) (*models.OnePortfolioResp, []*models.DealResp, error)
	GetPortfolioClosedDeals(ctx context.Context, portfolioID int) ([]*models.DealResp, error)
	CreatePortfolio(ctx context.Context, userID int, authType string, portfolio *models.Portfolio) (*models.Portfolio, error)
	DeletePortfolio(ctx context.Context, portfolioID int) error
}

type MarketRepository interface {
	GetAllStocks(ctx context.Context) ([]*models.DealResp, error)
	GetOneStock(ctx context.Context, stockID int) (*models.DealResp, error)
	CreateDeal(ctx context.Context, stockID int, stockAmount int, portfolioID int) (int, error)
}

type DealRepository interface {
	GetOneDeal(ctx context.Context, dealID int) (*models.DealResp, error)
	CloseDeal(ctx context.Context, dealID int) error
	DeleteDeal(ctx context.Context, dealID int) error
	GetOneClosedDeal(ctx context.Context, closedDealID int) (*models.DealResp, error)
	DeleteClosedDeal(ctx context.Context, closedDealID int) error
}
