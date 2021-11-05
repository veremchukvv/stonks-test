package pg

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/veremchukvv/stonks-test/internal/models"
	"github.com/veremchukvv/stonks-test/pkg/logging"
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

func (ur *PostgresUserRepo) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	log := logging.FromContext(ctx)
	const query string = `SELECT id, name, username, email, password_hash FROM users WHERE email=$1`
	var u models.User
	err := ur.db.QueryRow(ctx, query, email).Scan(&u.Id, &u.Name,
		&u.Username, &u.Email, &u.Password)
	if err != nil {
		log.Errorf("User not found %v", err)
		return nil, err
	}
	return &u, nil
}

func (ur *PostgresUserRepo) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	log := logging.FromContext(ctx)
	const query string = `INSERT INTO users 
		(name, username, email, password_hash) 
		VALUES($1, $2, $3, $4)
		returning id;`

	var uid int
	err := ur.db.QueryRow(ctx, query, user.Name, user.Username, user.Email, user.Password).Scan(&uid)
	if err != nil {
		log.Errorf("Error on write user to database %v", err)
		return nil, err
	}
	user.Id = uid
	return user, nil
}
