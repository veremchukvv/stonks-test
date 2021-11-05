package service

import (
	"context"
	"github.com/veremchukvv/stonks-test/internal/models"
	"github.com/veremchukvv/stonks-test/internal/repository"
	"github.com/veremchukvv/stonks-test/pkg/hash"
)

type UserSignupReq struct {
	Name     string
	Email    string
	Password string
}

type UserSigninReq struct {
	Email    string
	Password string
}

type UserServiceImp struct {
	repo   repository.UserRepository
	hasher hash.PasswordHasher
}

func NewUserServiceImp(repo repository.UserRepository, hasher *hash.BCHasher) *UserServiceImp {
	return &UserServiceImp{
		repo,
		hasher,
	}
}

func (us *UserServiceImp) GetUser(ctx context.Context, ID int) (*models.User, error) {
	return nil, nil
}

func (us *UserServiceImp) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	var err error
	user.Password, err = us.hasher.Hash(user.Password)
	if err != nil {
		return nil, err
	}
	return us.repo.CreateUser(ctx, user)
}

func (us *UserServiceImp) GenerateToken(userid string, password string) (string, error) {
	return "", nil
}

func (us *UserServiceImp) ParseToken(token string) (int, error) {
	return 0, nil
}
