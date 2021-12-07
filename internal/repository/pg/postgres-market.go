package pg

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/veremchukvv/stonks-test/internal/models"
	"github.com/veremchukvv/stonks-test/pkg/logging"
)

type PostgresMarketRepo struct {
	db  *pgxpool.Pool
	ctx context.Context
}

func NewPostgresMarketRepo(db *pgxpool.Pool, ctx context.Context) *PostgresMarketRepo {
	return &PostgresMarketRepo{
		db,
		ctx,
	}
}

func (pmr *PostgresMarketRepo) GetAllStocks(ctx context.Context) ([]*models.DealResp, error) {
	log := logging.FromContext(ctx)

	const queryAllStocks string = `SELECT stock_id, stock_name, ticker, stock_type, cost, currency_ticker
									FROM stocks INNER JOIN currencies ON currency_id = currency WHERE cost>0`
	var stocks []*models.DealResp
	rowsPortfolios, err := pmr.db.Query(ctx, queryAllStocks)
	if err != nil {
		log.Infof("Error on query rows: %v", err)
		return nil, err
	}
	defer rowsPortfolios.Close()
	for rowsPortfolios.Next() {
		var stock models.DealResp
		err = rowsPortfolios.Scan(&stock.Id, &stock.Name, &stock.Ticker, &stock.Type, &stock.Cost, &stock.Currency)
		if err != nil {
			log.Infof("Error on scan rows: %v", err)
			return nil, err
		}
		stocks = append(stocks, &stock)
	}
	return stocks, nil
}

func (pmr *PostgresMarketRepo) GetOneStock(ctx context.Context, stockId int) (*models.DealResp, error) {
	log := logging.FromContext(ctx)

	const query string = `SELECT stock_name, ticker, stock_type, description, cost, currency_ticker
									FROM stocks INNER JOIN currencies ON currency_id = currency WHERE stock_id=$1`
	var stock models.DealResp
	err := pmr.db.QueryRow(ctx, query, stockId).Scan(&stock.Name, &stock.Ticker, &stock.Type, &stock.Description,
		&stock.Cost, &stock.Currency)
	if err != nil {
		log.Infof("Error on query rows: %v", err)
		return nil, err
	}
	return &stock, nil
}
func (pmr *PostgresMarketRepo) CreateDeal(ctx context.Context, stockId int, stockAmount int, portfolioId int) (int, error) {
	log := logging.FromContext(ctx)

	const query string = `WITH rows AS (SELECT cost, currency FROM stocks WHERE stock_id = $1) INSERT INTO deals 
						(portfolio, stock_item, amount, stock_cost, stock_currency, buy_cost) SELECT $2, $3, $4, rows.cost, 
						rows.currency, cost FROM rows returning deal_id;`
	var did int
	err := pmr.db.QueryRow(ctx, query, stockId, portfolioId, stockId, stockAmount).Scan(&did)
	if err != nil {
		log.Infof("Error on query rows: %v", err)
		return 0, err
	}

	return did, nil
}