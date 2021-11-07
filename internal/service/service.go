package service

import (
	"context"
	"github.com/veremchukvv/stonks-test/internal/models"
	"github.com/veremchukvv/stonks-test/internal/repository"
	"github.com/veremchukvv/stonks-test/pkg/hash"
)

type UserService interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	GetUser(ctx context.Context, token string) (*models.User, *models.VKUser, error)
	CreateVKUser(ctx context.Context, user *models.VKUser) (*models.VKUser, error)
	GetVKUserByID(ctx context.Context, vkid int) (*models.VKUser, error)
	GenerateToken(ctx context.Context, email string, password string) (string, error)
	GenerateVKToken(ctx context.Context, vkid int) (string, error)
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
