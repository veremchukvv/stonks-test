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
	err = pr.db.QueryRow(ctx, queryNewStockItem, pid, 1, 0, 1, 1).Scan(&sid)
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

	currencyList, err := pr.getAllCurrencies(ctx)
	if err != nil {
		log.Infof("error on fetch currencies from DB %v", err)
		return nil, err
	}

	AssetsList := make([]string, 0, 5)
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

	for i, port := range portfolios {
		for _, cur := range currencyList {
			log.Infof("portfolio ID: %d", port.Id)
			log.Infof("currency ID: %d", cur.Id)
			const query string = `SELECT SUM(stock_value) FROM stocks_items WHERE (portfolio=$1 and stock_currency=$2)`

			err := pr.db.QueryRow(ctx, query, port.Id, cur.Id).Scan(&AssetsR.Results)
			if err != nil {
				log.Infof("Error on scan rows: %v", err)
				return nil, err
			}

			////PortPtr := reflect.ValueOf(&port).Elem()
			//log.Info(PortPtr)
			//log.Info(PortPtr.Kind())
			//println(PortPtr.String())
			//println(PortPtr.Elem().String())
			//log.Info(portPtr.FieldByName(assetsList[i]))
			//log.Info(portPtr.FieldByName(assetsList[i]).Elem())
			//indir := reflect.Indirect(PortPtr)
			//log.Info(indir)
			//log.Info(indir.Kind())
			//reflect.Indirect(port).FieldByName(AssetsList[i]).SetFloat(AssetsR.Results)
			reflect.ValueOf(port).Elem().FieldByName(AssetsList[i]).SetFloat(AssetsR.Results)
		log.Info(AssetsR.Results)
			//if err != nil {
			//	log.Infof("Error on query rows: %v", err)
			//	return nil, err
			//}
		}
	}
	return portfolios, nil
}

func (pr *PostgresPortfolioRepo) getAllCurrencies(ctx context.Context) ([]*models.Currency, error) {
	log := logging.FromContext(ctx)

	const query string = `SELECT currency_id, ticker FROM currencies WHERE currency_id > 0`
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
	log.Info(currencies)
	return currencies, nil
}
