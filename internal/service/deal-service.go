package service

import (
	"context"
	"github.com/golang-jwt/jwt"
	"github.com/veremchukvv/stonks-test/internal/models"
	"github.com/veremchukvv/stonks-test/internal/repository"
	"github.com/veremchukvv/stonks-test/pkg/logging"
)

type dealServiceImp struct {
	repo repository.DealRepository
}

func NewDealServiceImp(repo repository.DealRepository) *dealServiceImp {
	return &dealServiceImp{
		repo,
	}
}

func (dsi *dealServiceImp) GetOneDeal(ctx context.Context, token string, dealId int) (*models.DealResp, error) {
	log := logging.FromContext(ctx)

	_, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(key *jwt.Token) (interface{}, error) {
		return []byte(SignKey), nil
	})
	if err != nil {
		log.Info("error on authenticating user")
		return nil, err
	}

	return dsi.repo.GetOneDeal(ctx, dealId)
}

func (dsi *dealServiceImp) GetOneClosedDeal(ctx context.Context, token string, closedDealId int) (*models.DealResp, error) {
	log := logging.FromContext(ctx)

	_, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(key *jwt.Token) (interface{}, error) {
		return []byte(SignKey), nil
	})
	if err != nil {
		log.Info("error on authenticating user")
		return nil, err
	}

	return dsi.repo.GetOneClosedDeal(ctx, closedDealId)
}

func (dsi *dealServiceImp) CloseDeal(ctx context.Context, token string, dealId int) error {
	log := logging.FromContext(ctx)

	_, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(key *jwt.Token) (interface{}, error) {
		return []byte(SignKey), nil
	})
	if err != nil {
		log.Info("error on authenticating user")
		return err
	}

	return dsi.repo.CloseDeal(ctx, dealId)
}

func (dsi *dealServiceImp) DeleteDeal(ctx context.Context, token string, dealId int) error {
	log := logging.FromContext(ctx)

	_, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(key *jwt.Token) (interface{}, error) {
		return []byte(SignKey), nil
	})
	if err != nil {
		log.Info("error on authenticating user")
		return err
	}

	return dsi.repo.DeleteDeal(ctx, dealId)
}

func (dsi *dealServiceImp) DeleteClosedDeal(ctx context.Context, token string, closedDealId int) error {
	log := logging.FromContext(ctx)

	_, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(key *jwt.Token) (interface{}, error) {
		return []byte(SignKey), nil
	})
	if err != nil {
		log.Info("error on authenticating user")
		return err
	}

	return dsi.repo.DeleteClosedDeal(ctx, closedDealId)
}
