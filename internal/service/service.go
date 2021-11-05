package service

import (
	"context"
	"github.com/veremchukvv/stonks-test/internal/models"
	"github.com/veremchukvv/stonks-test/internal/repository"
	"github.com/veremchukvv/stonks-test/pkg/hash"
)

type UserService interface {
	GetUser(ctx context.Context, ID int) (*models.User, error)
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	GenerateToken(userid string, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Services struct {
	UserService UserService
}

func NewService(store *repository.Store, hasher *hash.BCHasher) *Services {
	return &Services{
		NewUserServiceImp(store, hasher),
	}
}
