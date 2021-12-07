package pg

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/veremchukvv/stonks-test/internal/models"
	"github.com/veremchukvv/stonks-test/pkg/logging"
	"reflect"
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
	const queryPortfolios string = `SELECT portfolio_id, portfolio_name, description, is_public FROM portfolios WHERE 
									(user_id=$1 and user_auth_type=$2)`
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
	if portfolios == nil {
		return nil, nil
	}
	portfoliosWithAssets, err := pr.getPortfolioAssets(ctx, portfolios)
	return portfoliosWithAssets, nil
}

func (pr *PostgresPortfolioRepo) GetPortfolioDeals(ctx context.Context, portfolioId int) (*models.OnePortfolioResp, []*models.DealResp, error) {
	log := logging.FromContext(ctx)

	const queryPortfolio string = `SELECT portfolio_name, description, is_public FROM portfolios WHERE portfolio_id=$1`
	var portfolio models.OnePortfolioResp

	err := pr.db.QueryRow(ctx, queryPortfolio, portfolioId).Scan(&portfolio.Name, &portfolio.Description, &portfolio.Public)
	if err != nil {
		log.Infof("Error on query rows: %v", err)
		return nil, nil, err
	}

	const queryDeals string = `SELECT deal_id, ticker, stock_name, stock_type, amount, stock_cost, stock_value, 
					currency_ticker, opened_at, income_money, income_percent FROM deals 
                    INNER JOIN stocks ON stock_id = stock_item AND stock_currency = currency AND stock_cost = cost 
                    INNER JOIN currencies ON currency_id = stock_currency WHERE (portfolio=$1 and stock_cost>0)`

	var deals []*models.DealResp
	rowsDeals, err := pr.db.Query(ctx, queryDeals, portfolioId)
	if err != nil {
		log.Infof("Error on query rows: %v", err)
		return nil, nil, err
	}
	defer rowsDeals.Close()
	for rowsDeals.Next() {
		var deal models.DealResp
		err = rowsDeals.Scan(&deal.Id, &deal.Ticker, &deal.Name, &deal.Type, &deal.Amount, &deal.Cost, &deal.Value, &deal.Currency, &deal.OpenedAt, &deal.Profit, &deal.Percent)
		deals = append(deals, &deal)
	}

	return &portfolio, deals, nil
}

func (pr *PostgresPortfolioRepo) GetPortfolioClosedDeals(ctx context.Context, portfolioId int) ([]*models.DealResp, error) {
	log := logging.FromContext(ctx)

	const queryClosedDeals string = `SELECT closed_deal_id, ticker, stock_name, stock_type, sell_cost, amount, 
					currency_ticker, closed_at, stock_value, income_money, income_percent FROM closed_deals 
                    INNER JOIN stocks ON stock_id = stock_item AND stock_currency = currency 
                    INNER JOIN currencies ON currency_id = stock_currency WHERE (portfolio=$1 and stock_cost >0)`

	var closedDeals []*models.DealResp
	rowsClosedDeals, err := pr.db.Query(ctx, queryClosedDeals, portfolioId)
	if err != nil {
		log.Infof("Error on query rows: %v", err)
		return nil, err
	}
	defer rowsClosedDeals.Close()
	for rowsClosedDeals.Next() {
		var closedDeal models.DealResp
		err = rowsClosedDeals.Scan(&closedDeal.Id, &closedDeal.Ticker, &closedDeal.Name, &closedDeal.Type, &closedDeal.SellCost, &closedDeal.Amount, &closedDeal.Currency, &closedDeal.ClosedAt, &closedDeal.Value, &closedDeal.Profit, &closedDeal.Percent)
		closedDeals = append(closedDeals, &closedDeal)
	}

	return closedDeals, nil
}

func (pr *PostgresPortfolioRepo) CreatePortfolio(ctx context.Context, userId int, authType string, newPortfolio *models.Portfolio) (*models.Portfolio, error) {
	log := logging.FromContext(ctx)

	currenciesList, err := pr.getAllCurrencies(ctx)
	if err != nil {
		log.Infof("error on fetch currencies from DB %v", err)
		return nil, err
	}

	const createNewPortfolio string = `INSERT INTO portfolios (user_id, user_auth_type, portfolio_name, description, 
									is_public) VALUES ($1, $2, $3, $4, $5) returning portfolio_id`

	var pid int
	err = pr.db.QueryRow(ctx, createNewPortfolio, userId, authType, newPortfolio.Name, newPortfolio.Description, newPortfolio.Public).Scan(&pid)
	if err != nil {
		log.Infof("Error on processing query to DB: %v", err)
		return nil, err
	}

	var bid int
	var sid int
	var csid int
	const createNewBalances string = `INSERT INTO balances (portfolio_id, currency_id, money_value) VALUES ($1, $2, $3) 
									returning balance_id`
	const createNullStockItem string = `INSERT INTO deals (portfolio, stock_item, stock_cost, stock_currency, 
									amount) VALUES ($1, $2, $3, $4, $5) returning deal_id`
	const createNullClosedStockItem string = `INSERT INTO closed_deals (portfolio, stock_item, stock_cost, stock_currency, 
									amount) VALUES ($1, $2, $3, $4, $5) returning closed_deal_id`
	for i, v := range currenciesList {
		err = pr.db.QueryRow(ctx, createNewBalances, pid, v.Id, 0).Scan(&bid)
		if err != nil {
			log.Infof("Error on processing query to DB: %v", err)
			return nil, err
		}
		err = pr.db.QueryRow(ctx, createNullStockItem, pid, i+1, 0, v.Id, 1).Scan(&sid)
		if err != nil {
			log.Infof("Error on processing query to DB: %v", err)
			return nil, err
		}
		err = pr.db.QueryRow(ctx, createNullClosedStockItem, pid, i+1, 0, v.Id, 1).Scan(&csid)
		if err != nil {
			log.Infof("Error on processing query to DB: %v", err)
			return nil, err
		}
	}

	return newPortfolio, nil
}

func (pr *PostgresPortfolioRepo) DeletePortfolio(ctx context.Context, portfolioId int) error {
	log := logging.FromContext(ctx)

	const query string = `DELETE FROM portfolios WHERE portfolio_id =$1 returning portfolio_id`
	var pid int
	err := pr.db.QueryRow(ctx, query, portfolioId).Scan(&pid)
	if err != nil {
		log.Infof("error on deleting portfolio: %d from database %v", portfolioId, err)
		return err
	}
	return nil
}

func (pr *PostgresPortfolioRepo) getPortfolioAssets(ctx context.Context, portfolios []*models.Portfolio) ([]*models.Portfolio, error) {
	log := logging.FromContext(ctx)

	currencyList, err := pr.getAllCurrencies(ctx)
	if err != nil {
		log.Infof("error on fetch currencies from DB %v", err)
		return nil, err
	}

	AssetsList := make([]string, 0)
	ProfitList := make([]string, 0)
	PercentList := make([]string, 0)
	t := reflect.TypeOf(&models.Portfolio{}).Elem()
	for x := 0; x < t.NumField(); x++ {
		field := t.Field(x)
		fieldTag := field.Tag.Get("ticker")
		if fieldTag != "" {
			AssetsList = append(AssetsList, "Assets"+fieldTag)
			ProfitList = append(ProfitList, "Profit"+fieldTag)
			PercentList = append(PercentList, "Percent"+fieldTag)
		}
	}

	type Assets struct {
		Results float64
	}
	var AssetsR Assets

	//TODO check if struct really needed here

	for _, port := range portfolios {
		for j, cur := range currencyList {
			const queryAssets string = `WITH rows AS (SELECT (SELECT SUM(stock_value) FROM deals WHERE 
                                 (portfolio=$1 and stock_currency=$2)) AS sum1, (SELECT SUM(stock_value) FROM closed_deals WHERE 
                                 portfolio=$1 and stock_currency=$2) AS sum2) SELECT SUM(sum1+sum2) from rows`
			err := pr.db.QueryRow(ctx, queryAssets, port.Id, cur.Id).Scan(&AssetsR.Results)
			if err != nil {
				log.Infof("Error on scan rows: %v", err)
				return nil, err
			}
			reflect.ValueOf(port).Elem().FieldByName(AssetsList[j]).SetFloat(AssetsR.Results)

			const queryProfit string = `select coalesce(sum(income_money), 0) FROM deals WHERE (portfolio=$1 and stock_currency=$2)`
			err = pr.db.QueryRow(ctx, queryProfit, port.Id, cur.Id).Scan(&AssetsR.Results)
			if err != nil {
				log.Infof("Error on scan rows: %v", err)
				return nil, err
			}
			reflect.ValueOf(port).Elem().FieldByName(ProfitList[j]).SetFloat(AssetsR.Results)

			const queryPercent string = `select coalesce(avg(income_percent), 0) FROM deals WHERE (portfolio=$1 and stock_currency=$2)`
			err = pr.db.QueryRow(ctx, queryPercent, port.Id, cur.Id).Scan(&AssetsR.Results)
			if err != nil {
				log.Infof("Error on scan rows: %v", err)
				return nil, err
			}
			reflect.ValueOf(port).Elem().FieldByName(PercentList[j]).SetFloat(AssetsR.Results)

		}
	}
	return portfolios, nil
}

func (pr *PostgresPortfolioRepo) getAllCurrencies(ctx context.Context) ([]*models.Currency, error) {
	log := logging.FromContext(ctx)

	const query string = `SELECT currency_id, currency_ticker FROM currencies WHERE currency_id > 0`
	var currencies []*models.Currency
	rows, err := pr.db.Query(ctx, query)
	if err != nil {
		log.Infof("Error on query rows: %v", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var currency models.Currency
		err = rows.Scan(&currency.Id, &currency.Ticker)
		if err != nil {
			log.Infof("Error on scan rows: %v", err)
			return nil, err
		}
		currencies = append(currencies, &currency)
	}
	return currencies, nil
}
