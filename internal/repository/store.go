package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/veremchukvv/stonks-test/internal/repository/pg"
)

type Store struct {
	UserRepository
	PortfolioRepository
	MarketRepository
}

func NewStore(db *pgxpool.Pool, ctx context.Context) *Store {
	return &Store{
		pg.NewPostgresUserRepo(db, ctx),
		pg.NewPostgresPortfolioRepo(db, ctx),
		pg.NewPostgresMarketRepo(db, ctx),
	}
}
