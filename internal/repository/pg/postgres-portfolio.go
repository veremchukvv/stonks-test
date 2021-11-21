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

	currenciesList, err := pr.getAllCurrencies(ctx)
	if err != nil {
		log.Infof("error on fetch currencies from DB %v", err)
		return nil, err
	}

	const createNewPortfolio string = `INSERT INTO portfolios (user_id, user_auth_type, portfolio_name, description, is_public) VALUES ($1, $2, $3, $4, $5) returning portfolio_id`

	var pid int
	err = pr.db.QueryRow(ctx, createNewPortfolio, userId, authType, newPortfolio.Name, newPortfolio.Description, newPortfolio.Public).Scan(&pid)
	if err != nil {
		log.Infof("Error on processing query to DB: %v", err)
		return nil, err
	}

	var bid int
	var sid int
	const createNewBalances string = `INSERT INTO balances (portfolio_id, currency_id, money_value) VALUES ($1, $2, $3) returning balance_id`
	const createNewStockItem string = `INSERT INTO stocks_items (portfolio, stock_item, stock_cost, stock_currency, amount) VALUES ($1, $2, $3, $4, $5) returning stocks_item_id`
	for _, v := range currenciesList {
		err = pr.db.QueryRow(ctx, createNewBalances, pid, v.Id, 0).Scan(&bid)
		if err != nil {
			log.Infof("Error on processing query to DB: %v", err)
			return nil, err
		}
		err = pr.db.QueryRow(ctx, createNewStockItem, pid, 1, 0, v.Id, 1).Scan(&sid)
		if err != nil {
			log.Infof("Error on processing query to DB: %v", err)
			return nil, err
		}
	}
	return newPortfolio, nil
}

func (pr *PostgresPortfolioRepo) DeletePortfolio(ctx context.Context, portfolioId int) error {
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
