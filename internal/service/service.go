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
	CreateVKUser(ctx context.Context, user *models.User) (*models.User, error)
	GetVKUserByID(ctx context.Context, vkid int) (*models.User, error)
	GenerateToken(ctx context.Context, email string, password string) (string, error)
	GenerateVKToken(ctx context.Context, vkid int) (string, error)
	ParseToken(token string) (int, error)
}

type PortfolioService interface {
	GetAllPortfolios(ctx context.Context, token string) ([]*models.Portfolio, error)
	GetOnePortfolio(ctx context.Context, portfolio_id int) (*models.Portfolio, error)
	CreatePortfolio(ctx context.Context, token string, portfolio *models.Portfolio) (*models.Portfolio, error)
	DeletePortfolio(ctx context.Context, portfolio_id int) error
}

type Services struct {
	UserService      UserService
	PortfolioService PortfolioService
}

func NewService(store *repository.Store, hasher *hash.BCHasher) *Services {
	return &Services{
		NewUserServiceImp(store, hasher),
		NewPortfolioServiceImp(store),
	}
}
