package repository

import (
	"context"
	"github.com/veremchukvv/stonks-test/internal/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	FindUser(ctx context.Context, user *models.User) (*models.User, error)
}