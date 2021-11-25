package pg

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/veremchukvv/stonks-test/internal/models"
)

type PostgresMarketRepo struct {
	db *pgxpool.Pool
	ctx context.Context
}

func NewPostgresMarketRepo(db *pgxpool.Pool, ctx context.Context) *PostgresMarketRepo {
	return &PostgresMarketRepo{
		db,
		ctx,
	}
}

func (pmr *PostgresMarketRepo) GetAllStocks(ctx context.Context) ([]*models.Stock, error) {
	return nil, nil
}

func (pmr *PostgresMarketRepo) GetOneStock(ctx context.Context, stockId int) (*models.Stock, error) {
	return nil, nil
}
func (pmr *PostgresMarketRepo) CreateDeal(ctx context.Context, stockId int, stockAmount int) (int, error) {
	return 0, nil
}
func (pmr *PostgresMarketRepo) DeleteDeal(ctx context.Context, dealId int) error {
	return nil
}