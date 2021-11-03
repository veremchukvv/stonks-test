package repo

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type PG struct {
	db *pgxpool.Pool
}

func NewPG(ctx context.Context, url string) (*PG, error) {
	cfg, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, err
	}
	pool, err := pgxpool.ConnectConfig(ctx, cfg)
	if err != nil {
		return nil, err
	}
return &PG{pool}, nil
}

func (pg *PG) Health(ctx context.Context) error {
	var greeting string
	return pg.db.QueryRow(ctx, "Select version();").Scan(&greeting)
}
