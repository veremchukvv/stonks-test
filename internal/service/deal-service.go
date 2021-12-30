package service

import (
	"context"

	"github.com/golang-jwt/jwt"
	"github.com/veremchukvv/stonks-test/internal/models"
	"github.com/veremchukvv/stonks-test/internal/repository"
	"github.com/veremchukvv/stonks-test/pkg/logging"
)

type DealServiceImp struct {
	repo repository.DealRepository
}

func NewDealServiceImp(repo repository.DealRepository) *DealServiceImp {
	return &DealServiceImp{
		repo,
	}
}

func (dsi *DealServiceImp) GetOneDeal(ctx context.Context, token string, dealID int) (*models.DealResp, error) {
	log := logging.FromContext(ctx)

	_, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(key *jwt.Token) (interface{}, error) {
		return []byte(SignKey), nil
	})
	if err != nil {
		log.Info("error on authenticating user")
		return nil, err
	}

	return dsi.repo.GetOneDeal(ctx, dealID)
}

func (dsi *DealServiceImp) GetOneClosedDeal(ctx context.Context, token string, closedDealID int) (*models.DealResp, error) {
	log := logging.FromContext(ctx)

	_, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(key *jwt.Token) (interface{}, error) {
		return []byte(SignKey), nil
	})
	if err != nil {
		log.Info("error on authenticating user")
		return nil, err
	}

	return dsi.repo.GetOneClosedDeal(ctx, closedDealID)
}

func (dsi *DealServiceImp) CloseDeal(ctx context.Context, token string, dealID int) error {
	log := logging.FromContext(ctx)

	_, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(key *jwt.Token) (interface{}, error) {
		return []byte(SignKey), nil
	})
	if err != nil {
		log.Info("error on authenticating user")
		return err
	}

	return dsi.repo.CloseDeal(ctx, dealID)
}

func (dsi *DealServiceImp) DeleteDeal(ctx context.Context, token string, dealID int) error {
	log := logging.FromContext(ctx)

	_, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(key *jwt.Token) (interface{}, error) {
		return []byte(SignKey), nil
	})
	if err != nil {
		log.Info("error on authenticating user")
		return err
	}

	return dsi.repo.DeleteDeal(ctx, dealID)
}

func (dsi *DealServiceImp) DeleteClosedDeal(ctx context.Context, token string, closedDealID int) error {
	log := logging.FromContext(ctx)

	_, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(key *jwt.Token) (interface{}, error) {
		return []byte(SignKey), nil
	})
	if err != nil {
		log.Info("error on authenticating user")
		return err
	}

	return dsi.repo.DeleteClosedDeal(ctx, closedDealID)
}
