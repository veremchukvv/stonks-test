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

	portfolios, err := ps.repo.GetAllPortfolios(ctx, claims.UserId, claims.AuthType)

	_, err = json.Marshal(portfolios)
	if err != nil {
		log.Infof("error on marshalling portfolios %v", err)
		return nil, err
	}

	return portfolios, nil
}

func (ps *PortfolioServiceImp) GetOnePortfolio(ctx context.Context, portfolioId int) (*models.Portfolio, error) {
	return nil, nil
}

func (ps *PortfolioServiceImp) CreatePortfolio(ctx context.Context, token string, newPortfolio *models.Portfolio) (*models.Portfolio, error) {
	log := logging.FromContext(ctx)

	parsedToken, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(key *jwt.Token) (interface{}, error) {
		return []byte(SignKey), nil
	})
	if err != nil {
		return nil, err
	}
	//TODO move parse of jwt to middleware or func
	claims := parsedToken.Claims.(*tokenClaims)

	createdPortfolio, err := ps.repo.CreatePortfolio(ctx, claims.UserId, claims.AuthType, newPortfolio)
	if err != nil {
		log.Infof("Error on creating portfolio in DB %v", err)
		return nil, err
	}
	return createdPortfolio, nil
}

func (ps *PortfolioServiceImp) DeletePortfolio(ctx context.Context, portfolioId int) error {
	return nil
}
