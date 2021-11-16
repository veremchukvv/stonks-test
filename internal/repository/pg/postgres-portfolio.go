package pg

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/veremchukvv/stonks-test/internal/models"
	"github.com/veremchukvv/stonks-test/pkg/logging"
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

func (pr *PostgresPortfolioRepo) GetAllPortfolios(ctx context.Context, userId int, authType string) ([]*models.Portfolio, error) {
	log := logging.FromContext(ctx)
	const query string = `SELECT portfolio_id, portfolio_name, description, is_public FROM portfolios WHERE (user_id=$1 and user_auth_type=$2)`
	var portfolios []*models.Portfolio
	rows, err := pr.db.Query(ctx, query, userId, authType)
	if err != nil {
		log.Infof("Error on query rows: %v", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var portfolio models.Portfolio
		err = rows.Scan(&portfolio.Id, &portfolio.Name, &portfolio.Description, &portfolio.Public)
		if err != nil {
			log.Infof("Error on scan rows: %v", err)
			return nil, err
		}
		portfolios = append(portfolios, &portfolio)
	}

	return portfolios, nil
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
