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
	GetVKUserByID(ctx context.Context, id int) (*models.User, error)
	CreateGoogleUser(ctx context.Context, user *models.User) (*models.User, error)
	GetGoogleUserByID(ctx context.Context, gid string) (*models.User, error)
	GenerateToken(ctx context.Context, email string, password string) (string, error)
	GenerateVKToken(id int) (string, error)
	GenerateGoogleToken(id int) (string, error)
	//ParseToken(token string) (int, error)
}

type PortfolioService interface {
	GetAllPortfolios(ctx context.Context, token string) ([]*models.Portfolio, error)
	GetPortfolioDeals(ctx context.Context, token string, portfolioId int) (*models.OnePortfolioResp, []*models.DealResp, error)
	GetPortfolioClosedDeals(ctx context.Context, token string, portfolioId int) ([]*models.DealResp, error)
	CreatePortfolio(ctx context.Context, token string, portfolio *models.Portfolio) (*models.Portfolio, error)
	DeletePortfolio(ctx context.Context, token string, portfolioId int) error
}

type MarketService interface {
	GetAllStocks(ctx context.Context) ([]*models.DealResp, error)
	GetOneStock(ctx context.Context, stockId int) (*models.DealResp, error)
	CreateDeal(ctx context.Context, token string, stockId int, stockAmount int, portfolioId int) (int, error)
}

type DealService interface {
	GetOneDeal(ctx context.Context, token string, dealId int) (*models.DealResp, error)
	CloseDeal(ctx context.Context, token string, dealId int) error
	DeleteDeal(ctx context.Context, token string, dealId int) error
	GetOneClosedDeal(ctx context.Context, token string, closedDealId int) (*models.DealResp, error)
	DeleteClosedDeal(ctx context.Context, token string, closedDealId int) error
}

type Services struct {
	UserService      UserService
	PortfolioService PortfolioService
	MarketService    MarketService
	DealService      DealService
}

func NewService(store *repository.Store, hasher *hash.BCHasher) *Services {
	return &Services{
		NewUserServiceImp(store, hasher),
		NewPortfolioServiceImp(store),
		NewMarketServiceImp(store),
		NewDealServiceImp(store),
	}
}
