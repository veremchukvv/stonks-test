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
	const queryPortfolios string = `SELECT portfolio_id, portfolio_name, description, is_public FROM portfolios WHERE (user_id=$1 and user_auth_type=$2)`
	var portfolios []*models.Portfolio
	rowsPortfolios, err := pr.db.Query(ctx, queryPortfolios, userId, authType)
	if err != nil {
		log.Infof("Error on query rows: %v", err)
		return nil, err
	}
	defer rowsPortfolios.Close()
	for rowsPortfolios.Next() {
		var portfolio models.Portfolio
		err = rowsPortfolios.Scan(&portfolio.Id, &portfolio.Name, &portfolio.Description, &portfolio.Public)
		if err != nil {
			log.Infof("Error on scan rows: %v", err)
			return nil, err
		}
		portfolios = append(portfolios, &portfolio)
	}
	portfoliosWithAssets, err := pr.getPortfolioAssets(ctx, portfolios)
	return portfoliosWithAssets, nil
}

func (pr *PostgresPortfolioRepo) GetOnePortfolio(ctx context.Context, portfolioId int) (*models.Portfolio, error) {
	return nil, nil
}

func (pr *PostgresPortfolioRepo) CreatePortfolio(ctx context.Context, userId int, authType string, newPortfolio *models.Portfolio) (*models.Portfolio, error) {
	log := logging.FromContext(ctx)

	const queryNewPortfolio string = `INSERT INTO portfolios (user_id, user_auth_type, portfolio_name, description, is_public) VALUES ($1, $2, $3, $4, $5)`
	var portfolio *models.Portfolio
	var pid int
	err := pr.db.QueryRow(ctx, queryNewPortfolio, userId, authType, newPortfolio.Name, newPortfolio.Description, newPortfolio.Public).Scan(&pid)
	if err != nil {
		log.Infof("Error on processing query to DB: %v", err)
		return nil, err
	}
	var bid int
	const queryNewBalances string = `INSERT INTO balances (portfolio_id, currency_id, money_value) VALUES ($1, $2, $3)`
	err = pr.db.QueryRow(ctx, queryNewBalances, pid, 1, 0).Scan(&bid)
	if err != nil {
		log.Infof("Error on processing query to DB: %v", err)
		return nil, err
	}
	err = pr.db.QueryRow(ctx, queryNewBalances, pid, 2, 0).Scan(&bid)
	if err != nil {
		log.Infof("Error on processing query to DB: %v", err)
		return nil, err
	}
	err = pr.db.QueryRow(ctx, queryNewBalances, pid, 3, 0).Scan(&bid)
	if err != nil {
		log.Infof("Error on processing query to DB: %v", err)
		return nil, err
	}
	var sid int
	const queryNewStockItem string = `INSERT INTO stock_items (portfolio, stock_item, stock_cost, stock_currency, amount) VALUES ($1, $2, $3, $4, $5)`
	err = pr.db.QueryRow(ctx, queryNewStockItem, pid, 7, 0, 1, 1).Scan(&sid)
	if err != nil {
		log.Infof("Error on processing query to DB: %v", err)
		return nil, err
	}
	portfolio.Id = pid
	return portfolio, nil
}

func (pr *PostgresPortfolioRepo) DeletePortfolio(ctx context.Context, portfolioId int) error {
	return nil
}

func (pr *PostgresPortfolioRepo) getPortfolioAssets(ctx context.Context, portfolios []*models.Portfolio) ([]*models.Portfolio, error) {
	log := logging.FromContext(ctx)

	for i, _ := range portfolios {
		const query string = `SELECT SUM(stock_value) FROM stocks_items WHERE (portfolio=$1 and stock_currency=1)`
		err := pr.db.QueryRow(ctx, query, portfolios[i].Id).Scan(&portfolios[i].AssetsRUB)
		if err != nil {
			log.Infof("Error on query rows: %v", err)
			return nil, err
		}
	}
	return portfolios, nil
}
