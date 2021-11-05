package pg

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/veremchukvv/stonks-test/internal/models"
)

type PostgresUserRepo struct {
	db  *pgxpool.Pool
	ctx context.Context
}

func NewPostgresUserRepo(pgpool *pgxpool.Pool, ctx context.Context) *PostgresUserRepo {
	return &PostgresUserRepo{
		pgpool,
		ctx}
}

func (ur *PostgresUserRepo) FindUser(ctx context.Context, user *models.User) (*models.User, error) {
	return nil, nil
}

func (ur *PostgresUserRepo) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	ur.db.QueryRow(ctx, )
	return nil, nil
}
