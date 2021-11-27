package service

import (
	"context"
	"github.com/veremchukvv/stonks-test/internal/models"
	"github.com/veremchukvv/stonks-test/internal/repository"
	"github.com/veremchukvv/stonks-test/pkg/hash"
)

type UserService interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	GetUser(ctx context.Context, token string) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User, token string) (*models.User, error)
	DeleteUser(ctx context.Context, token string) error
	CreateVKUser(ctx context.Context, user *models.User) (*models.User, error)
	GetVKUserByID(ctx context.Context, vkid int) (*models.User, error)
	GenerateToken(ctx context.Context, email string, password string) (string, error)
	GenerateVKToken(ctx context.Context, vkid int) (string, error)
	ParseToken(token string) (int, error)
}

type PortfolioService interface {
	GetAllPortfolios(ctx context.Context, token string) ([]*models.Portfolio, error)
	GetOnePortfolio(ctx context.Context, token string, portfolioId int) (*models.OnePortfolioResp, []*models.StockResp, error)
	CreatePortfolio(ctx context.Context, token string, portfolio *models.Portfolio) (*models.Portfolio, error)
	DeletePortfolio(ctx context.Context, token string, portfolioId int) error
}

type MarketService interface {
	GetAllStocks(ctx context.Context) ([]*models.StockResp, error)
	GetOneStock(ctx context.Context, stockId int) (*models.Stock, error)
	CreateDeal(ctx context.Context, token string, stockId int, stockAmount int) (int, error)
	DeleteDeal(ctx context.Context, token string, dealId int) error
}

type Services struct {
	UserService      UserService
	PortfolioService PortfolioService
	MarketService    MarketService
}

func NewService(store *repository.Store, hasher *hash.BCHasher) *Services {
	return &Services{
		NewUserServiceImp(store, hasher),
		NewPortfolioServiceImp(store),
		NewMarketServiceImp(store),
	}
}
