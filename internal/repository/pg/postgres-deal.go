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

func NewPostgresDealRepo(ctx context.Context, pgpool *pgxpool.Pool) *PostgresDealRepo {
	return &PostgresDealRepo{
		pgpool,
		ctx,
	}
}

func (npr *PostgresDealRepo) GetOneDeal(ctx context.Context, dealID int) (*models.DealResp, error) {
	log := logging.FromContext(ctx)

	const query string = `SELECT deal_id, ticker, stock_name, stock_type, description, buy_cost, income_money, income_percent, amount, stock_cost, stock_value, 
					currency_ticker, opened_at FROM deals INNER JOIN stocks ON stock_id = stock_item AND 
                    stock_currency = currency AND stock_cost = cost INNER JOIN currencies ON currency_id = 
                    stock_currency WHERE deal_id=$1`

	var deal models.DealResp

	err := npr.db.QueryRow(ctx, query, dealID).Scan(&deal.Id, &deal.Ticker, &deal.Name, &deal.Type, &deal.Description, &deal.BuyCost, &deal.Profit, &deal.Percent, &deal.Amount, &deal.Cost, &deal.Value, &deal.Currency, &deal.OpenedAt)
	if err != nil {
		log.Infof("Error on query rows: %v", err)
		return nil, err
	}

	return &deal, nil
}

func (npr *PostgresDealRepo) GetOneClosedDeal(ctx context.Context, closedDealID int) (*models.DealResp, error) {
	log := logging.FromContext(ctx)

	const query string = `SELECT closed_deal_id, stock_name, ticker, stock_type, description, amount, buy_cost, sell_cost, stock_value, 
					currency_ticker, opened_at, closed_at, income_money, income_percent FROM closed_deals INNER JOIN stocks ON stock_id = stock_item AND 
                    stock_currency = currency AND stock_cost = cost INNER JOIN currencies ON currency_id = 
                    stock_currency WHERE closed_deal_id=$1`

	var deal models.DealResp

	err := npr.db.QueryRow(ctx, query, closedDealID).Scan(&deal.Id, &deal.Name, &deal.Ticker, &deal.Type, &deal.Description, &deal.Amount, &deal.BuyCost, &deal.SellCost, &deal.Value, &deal.Currency, &deal.OpenedAt, &deal.ClosedAt, &deal.Profit, &deal.Percent)
	if err != nil {
		log.Infof("Error on query rows: %v", err)
		return nil, err
	}

	return &deal, nil
}

func (npr *PostgresDealRepo) CloseDeal(ctx context.Context, dealID int) error {
	log := logging.FromContext(ctx)

	const query string = `WITH preparation_table AS (SELECT deal_id, portfolio, stock_item, stock_cost, stock_currency, 
                          amount, opened_at, buy_cost, income_money, income_percent FROM deals WHERE deal_id = $1) 
                          INSERT INTO closed_deals (portfolio, stock_item, stock_cost, stock_currency, amount, opened_at, 
                          closed_at, buy_cost, sell_cost) SELECT portfolio, stock_item, stock_cost, stock_currency, 
                          amount, opened_at, NOW(), buy_cost, stock_cost FROM preparation_table WHERE deal_id = $1 
                          returning closed_deal_id`

	var cdid int
	err := npr.db.QueryRow(ctx, query, dealID).Scan(&cdid)
	if err != nil {
		log.Infof("error on closing deal: %d in database %v", dealID, err)
		return err
	}

	err = npr.DeleteDeal(ctx, dealID)
	if err != nil {
		log.Infof("error on closing deal (deleting opened deal): %d in database %v", dealID, err)
		return err
	}

	return nil
}

func (npr *PostgresDealRepo) DeleteDeal(ctx context.Context, dealID int) error {
	log := logging.FromContext(ctx)

	const query string = `DELETE FROM deals WHERE deal_id =$1 returning deal_id`
	var did int
	err := npr.db.QueryRow(ctx, query, dealID).Scan(&did)
	if err != nil {
		log.Infof("error on deleting stock item: %d from database %v", dealID, err)
		return err
	}
	return nil
}

func (npr *PostgresDealRepo) DeleteClosedDeal(ctx context.Context, closedDealID int) error {
	log := logging.FromContext(ctx)

	const query string = `DELETE FROM closed_deals WHERE closed_deal_id =$1 returning closed_deal_id`
	var did int
	err := npr.db.QueryRow(ctx, query, closedDealID).Scan(&did)
	if err != nil {
		log.Infof("error on deleting stock item: %d from database %v", closedDealID, err)
		return err
	}
	return nil
}
