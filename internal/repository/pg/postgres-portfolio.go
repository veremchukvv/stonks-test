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

func (pr *PostgresPortfolioRepo) GetOnePortfolio(ctx context.Context, portfolioId int) (*models.OnePortfolioResp, []*models.StockResp, error) {
	log := logging.FromContext(ctx)

	const queryPortfolio string = `SELECT portfolio_name, description, is_public FROM portfolios WHERE portfolio_id=$1`
	var portfolio models.OnePortfolioResp

	err := pr.db.QueryRow(ctx, queryPortfolio, portfolioId).Scan(&portfolio.Name, &portfolio.Description, &portfolio.Public)
	if err != nil {
		log.Infof("Error on query rows: %v", err)
		return nil, nil, err
	}

	const queryStocks string = `SELECT stocks_item_id, ticker, stock_name, stock_type, amount, stock_cost, stock_value, 
					currency_ticker, created_at, closed, income_final_money, income_final_percent FROM stocks_items 
                    INNER JOIN stocks ON stock_id = stock_item AND stock_currency = currency AND stock_cost = cost 
                    INNER JOIN currencies ON currency_id = stock_currency WHERE (portfolio=$1 and stock_cost>0)`
	var stocks []*models.StockResp
	rowsStocks, err := pr.db.Query(ctx, queryStocks, portfolioId)
	if err != nil {
		log.Infof("Error on query rows: %v", err)
		return nil, nil, err
	}
	defer rowsStocks.Close()
	for rowsStocks.Next() {
		var stock models.StockResp
		err = rowsStocks.Scan(&stock.Id, &stock.Ticker, &stock.Name, &stock.Type, &stock.Amount, &stock.Cost, &stock.Value, &stock.Currency, &stock.CreatedAt, &stock.IsClosed, &stock.ProfitClosed, &stock.PercentClosed)
		stocks = append(stocks, &stock)
	}
	return &portfolio, stocks, nil
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
	const createNewBalances string = `INSERT INTO balances (portfolio_id, currency_id, money_value) VALUES ($1, $2, $3) 
									returning balance_id`
	const createNewStockItem string = `INSERT INTO stocks_items (portfolio, stock_item, stock_cost, stock_currency, 
									amount) VALUES ($1, $2, $3, $4, $5) returning stocks_item_id`
	for i, v := range currenciesList {
		err = pr.db.QueryRow(ctx, createNewBalances, pid, v.Id, 0).Scan(&bid)
		if err != nil {
			log.Infof("Error on processing query to DB: %v", err)
			return nil, err
		}
		err = pr.db.QueryRow(ctx, createNewStockItem, pid, i+1, 0, v.Id, 1).Scan(&sid)
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
	t := reflect.TypeOf(&models.Portfolio{}).Elem()
	for x := 0; x < t.NumField(); x++ {
		field := t.Field(x)
		fieldTag := field.Tag.Get("ticker")
		if fieldTag != "" {
			AssetsList = append(AssetsList, "Assets"+fieldTag)
		}
	}

	type Assets struct {
		Results float64
	}
	var AssetsR Assets

	//TODO check if struct really needed here

	for _, port := range portfolios {
		for j, cur := range currencyList {
			const query string = `SELECT SUM(stock_value) FROM stocks_items WHERE (portfolio=$1 and stock_currency=$2)`
			err := pr.db.QueryRow(ctx, query, port.Id, cur.Id).Scan(&AssetsR.Results)
			if err != nil {
				log.Infof("Error on scan rows: %v", err)
				return nil, err
			}
			reflect.ValueOf(port).Elem().FieldByName(AssetsList[j]).SetFloat(AssetsR.Results)
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
