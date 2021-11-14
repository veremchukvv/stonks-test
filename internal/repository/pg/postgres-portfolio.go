package pg

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/veremchukvv/stonks-test/internal/models"
)

type PostgresPortfolioRepo struct {
	db  *pgxpool.Pool
	ctx context.Context
}

func NewPostgresPortfolioRepo(pgpool *pgxpool.Pool, ctx context.Context) *PostgresPortfolioRepo {
	return &PostgresPortfolioRepo{
		pgpool,
		ctx,
	}
}

func (pr *PostgresPortfolioRepo) GetAllPortfolios(ctx context.Context) ([]*models.Portfolio, error) {
	return nil, nil
}

func (pr *PostgresPortfolioRepo) GetOnePortfolio(ctx context.Context, portfolio_id int) (*models.Portfolio, error) {
	return nil, nil
}

func (pr *PostgresPortfolioRepo) CreatePortfolio(ctx context.Context, portfolio *models.Portfolio) (*models.Portfolio, error) {
	return nil, nil
}

func (pr *PostgresPortfolioRepo) DeletePortfolio(ctx context.Context, portfolio_id int) error {
	return nil
}

