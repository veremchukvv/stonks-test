package service

import (
	"context"
	"encoding/json"

	"github.com/golang-jwt/jwt"
	"github.com/veremchukvv/stonks-test/internal/models"
	"github.com/veremchukvv/stonks-test/internal/repository"
	"github.com/veremchukvv/stonks-test/pkg/logging"
)

type PortfolioServiceImp struct {
	repo repository.PortfolioRepository
}

func NewPortfolioServiceImp(repo repository.PortfolioRepository) *PortfolioServiceImp {
	return &PortfolioServiceImp{repo}
}

func (ps *PortfolioServiceImp) GetAllPortfolios(ctx context.Context, token string) ([]*models.Portfolio, error) {
	log := logging.FromContext(ctx)

	parsedToken, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(key *jwt.Token) (interface{}, error) {
		return []byte(SignKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims := parsedToken.Claims.(*tokenClaims)

	portfolios, err := ps.repo.GetAllPortfolios(ctx, claims.UserID, claims.AuthType)
	if err != nil {
		log.Infof("Error on query portfolios: %v", err)
		return nil, err
	}

	_, err = json.Marshal(portfolios)
	if err != nil {
		log.Infof("error on marshalling portfolios %v", err)
		return nil, err
	}

	return portfolios, nil
}

func (ps *PortfolioServiceImp) GetPortfolioDeals(ctx context.Context, token string, portfolioID int) (*models.OnePortfolioResp, []*models.DealResp, error) {
	log := logging.FromContext(ctx)

	_, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(key *jwt.Token) (interface{}, error) {
		return []byte(SignKey), nil
	})
	if err != nil {
		log.Info("error on authenticating user")
		return nil, nil, err
	}
	portfolio, deals, err := ps.repo.GetPortfolioDeals(ctx, portfolioID)
	if err != nil {
		log.Infof("error on fetching portfolio data from DB: %v", err)
		return nil, nil, err
	}
	return portfolio, deals, nil
}

func (ps *PortfolioServiceImp) GetPortfolioClosedDeals(ctx context.Context, token string, portfolioID int) ([]*models.DealResp, error) {
	log := logging.FromContext(ctx)

	_, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(key *jwt.Token) (interface{}, error) {
		return []byte(SignKey), nil
	})
	if err != nil {
		log.Info("error on authenticating user")
		return nil, err
	}
	closedDeals, err := ps.repo.GetPortfolioClosedDeals(ctx, portfolioID)
	if err != nil {
		log.Infof("error on fetching portfolio data from DB: %v", err)
		return nil, err
	}
	return closedDeals, nil
}

func (ps *PortfolioServiceImp) CreatePortfolio(ctx context.Context, token string, newPortfolio *models.Portfolio) (*models.Portfolio, error) {
	log := logging.FromContext(ctx)

	parsedToken, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(key *jwt.Token) (interface{}, error) {
		return []byte(SignKey), nil
	})
	if err != nil {
		log.Info("error on authenticating user")
		return nil, err
	}
	// TODO move parse of jwt to middleware or func
	claims := parsedToken.Claims.(*tokenClaims)

	createdPortfolio, err := ps.repo.CreatePortfolio(ctx, claims.UserID, claims.AuthType, newPortfolio)
	if err != nil {
		log.Infof("Error on creating portfolio in DB %v", err)
		return nil, err
	}
	return createdPortfolio, nil
}

func (ps *PortfolioServiceImp) DeletePortfolio(ctx context.Context, token string, portfolioID int) error {
	log := logging.FromContext(ctx)

	_, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(key *jwt.Token) (interface{}, error) {
		return []byte(SignKey), nil
	})
	if err != nil {
		log.Info("error on authenticating user")
		return err
	}

	err = ps.repo.DeletePortfolio(ctx, portfolioID)
	if err != nil {
		log.Infof("error on deleting portfolio %d", portfolioID)
		return err
	}

	return nil
}
