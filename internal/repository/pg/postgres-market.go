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

func (pmr *PostgresMarketRepo) GetAllStocks(ctx context.Context) ([]*models.StockResp, error) {
	log := logging.FromContext(ctx)

	const queryAllStocks string = `SELECT stock_id, stock_name, ticker, stock_type, cost, currency_ticker
									FROM stocks INNER JOIN currencies ON currency_id = currency WHERE cost>0`
	var stocks []*models.StockResp
	rowsPortfolios, err := pmr.db.Query(ctx, queryAllStocks)
	if err != nil {
		log.Infof("Error on query rows: %v", err)
		return nil, err
	}
	defer rowsPortfolios.Close()
	for rowsPortfolios.Next() {
		var stock models.StockResp
		err = rowsPortfolios.Scan(&stock.Id, &stock.Name, &stock.Ticker, &stock.Type, &stock.Cost, &stock.Currency)
		if err != nil {
			log.Infof("Error on scan rows: %v", err)
			return nil, err
		}
		stocks = append(stocks, &stock)
	}
	return stocks, nil
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
