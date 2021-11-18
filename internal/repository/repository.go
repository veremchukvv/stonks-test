package repository

import (
	"context"
	"github.com/veremchukvv/stonks-test/internal/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	CreateVKUser(ctx context.Context, user *models.User) (*models.User, error)
	GetVKUserByID(ctx context.Context, vkid int) (*models.User, error)
	GetUserByID(ctx context.Context, id int, authType string) (*models.User, error)
}

type PortfolioRepository interface {
	GetAllPortfolios(ctx context.Context, userId int, authType string) ([]*models.Portfolio, error)
	GetOnePortfolio(ctx context.Context, portfolio_id int) (*models.Portfolio, error)
	CreatePortfolio(ctx context.Context, userId int, authType string, portfolio *models.Portfolio) (*models.Portfolio, error)
	DeletePortfolio(ctx context.Context, portfolio_id int) error
}
