package pg

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/veremchukvv/stonks-test/internal/models"
	"github.com/veremchukvv/stonks-test/pkg/logging"
)

type PostgresDealRepo struct {
	db  *pgxpool.Pool
	ctx context.Context
}

func NewPostgresDealRepo(pgpool *pgxpool.Pool, ctx context.Context) *PostgresDealRepo {
	return &PostgresDealRepo{
		pgpool,
		ctx,
	}
}

func (npr *PostgresDealRepo) GetOneDeal(ctx context.Context, dealId int) (*models.StockResp, error) {
	log := logging.FromContext(ctx)

	const query string = `SELECT ticker, stock_name, stock_type, amount, stock_cost, stock_value, 
					currency_ticker, created_at FROM stocks_items INNER JOIN stocks ON stock_id = stock_item AND 
                    stock_currency = currency AND stock_cost = cost INNER JOIN currencies ON currency_id = 
                    stock_currency WHERE stocks_item_id=$1`

	var deal models.StockResp

	err := npr.db.QueryRow(ctx, query, dealId).Scan(&deal.Ticker, &deal.Name, &deal.Type, &deal.Amount, &deal.Cost, &deal.Value, &deal.Currency, &deal.CreatedAt)
	if err != nil {
		log.Infof("Error on query rows: %v", err)
		return nil, err
	}

	return &deal, nil

}
func (npr *PostgresDealRepo) CloseDeal(ctx context.Context, dealId int) error {
	log := logging.FromContext(ctx)

	const query string = `WITH aggregation_table AS (SELECT SUM((stock_cost - buy_cost) * amount) AS aggr_money, 
                          SUM((((stock_cost - buy_cost) * amount)/(buy_cost*amount)) * 100) AS aggr_percent FROM 
                          stocks_items WHERE stocks_item_id = $1) UPDATE stocks_items SET (closed, sell_cost, closed_at, 
                          income_final_money, income_final_percent) = (true, stock_cost, NOW(), 
                          aggregation_table.aggr_money, aggregation_table.aggr_percent) FROM aggregation_table WHERE 
                          stocks_item_id = $1 returning stocks_item_id`

	var did int
	err := npr.db.QueryRow(ctx, query, dealId).Scan(&did)
	if err != nil {
		log.Infof("error on closing deal: %d in database %v", dealId, err)
		return err
	}
	return nil
}
func (npr *PostgresDealRepo) DeleteDeal(ctx context.Context, dealId int) error {
	log := logging.FromContext(ctx)

	const query string = `DELETE FROM stocks_items WHERE stocks_item_id =$1 returning stocks_item_id`
	var did int
	err := npr.db.QueryRow(ctx, query, dealId).Scan(&did)
	if err != nil {
		log.Infof("error on deleting stock item: %d from database %v", dealId, err)
		return err
	}
	return nil
}