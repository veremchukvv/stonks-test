package pg

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type PG struct{}

func NewPG(ctx context.Context, url string) (*pgxpool.Pool, error) {
	cfg, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, err
	}
	pool, err := pgxpool.ConnectConfig(ctx, cfg)
	if err != nil {
		return nil, err
	}
	return pool, nil
}

func HealthPG(ctx context.Context, pool *pgxpool.Pool) error {
	var greeting string
	return pool.QueryRow(ctx, "Select version();").Scan(&greeting)
}
