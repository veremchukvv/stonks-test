package service

import (
	"context"
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"github.com/veremchukvv/stonks-test/internal/models"
	"github.com/veremchukvv/stonks-test/internal/repository"
	"github.com/veremchukvv/stonks-test/pkg/hash"
	"github.com/veremchukvv/stonks-test/pkg/logging"
	"os"
	"time"
)

const (
	tokenTTL = 12 * time.Hour
	//TODO move to config
)

var SignKey = os.Getenv("SIGN_KEY")

type tokenClaims struct {
	jwt.StandardClaims
	UserId   int    `json:"user_id"`
	AuthType string `json:"auth_type"`
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

func (us *UserServiceImp) GetUser(ctx context.Context, token string) (*models.User, error) {
	log := logging.FromContext(ctx)

	GetUserErr := errors.New("Can't get user")

	parsedToken, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(key *jwt.Token) (interface{}, error) {
		return []byte(SignKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims := parsedToken.Claims.(*tokenClaims)

	if claims.UserId != 0 {
		u, err := us.repo.GetUserByID(ctx, claims.UserId, claims.AuthType)
		if err != nil {
			log.Errorf("can't get user from db: %v", err)
			return nil, err
		}
		return u, nil
	}
	return nil, GetUserErr
}

func (us *UserServiceImp) GetVKUserByID(ctx context.Context, id int) (*models.User, error) {
	return us.repo.GetVKUserByID(ctx, id)
}

func (us *UserServiceImp) GetGoogleUserByID(ctx context.Context, gid string) (*models.User, error) {
	return us.repo.GetGoogleUserByID(ctx, gid)
}

func (us *UserServiceImp) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	var err error
	user.Password, err = us.hasher.Hash(user.Password)
	if err != nil {
		return nil, err
	}
	return us.repo.CreateUser(ctx, user)
}

func (us *UserServiceImp) UpdateUser(ctx context.Context, user *models.User, token string) (*models.User, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(key *jwt.Token) (interface{}, error) {
		return []byte(SignKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims := parsedToken.Claims.(*tokenClaims)
	user.Id = claims.UserId
	user.AuthType = claims.AuthType

	return us.repo.UpdateUser(ctx, user)
}

func (us *UserServiceImp) DeleteUser(ctx context.Context, token string) error {
	parsedToken, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(key *jwt.Token) (interface{}, error) {
		return []byte(SignKey), nil
	})
	if err != nil {
		return err
	}

	claims := parsedToken.Claims.(*tokenClaims)

	return us.repo.DeleteUser(ctx, claims.UserId, claims.AuthType)
}

func (us *UserServiceImp) CreateVKUser(ctx context.Context, user *models.User) (*models.User, error) {
	return us.repo.CreateVKUser(ctx, user)
}

func (us *UserServiceImp) CreateGoogleUser(ctx context.Context, user *models.User) (*models.User, error) {
	return us.repo.CreateGoogleUser(ctx, user)
}

func (us *UserServiceImp) GenerateToken(ctx context.Context, email string, password string) (string, error) {
	log := logging.FromContext(ctx)
	u, err := us.repo.GetUserByEmail(ctx, email)
	if err != nil {
		log.Info("Error on fetching user from DB")
		return "", err
	}
	hashedPassword := u.Password
	chk, err := us.hasher.CheckPWD(password, hashedPassword)
	if err != nil {
		return "", err
	}
	if chk == false {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		u.Id,
		"local",
	})
	return token.SignedString([]byte(SignKey))
}

func (us *UserServiceImp) GenerateVKToken(ctx context.Context, id int) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "stonks",
		},
		id,
		"vk",
	})
	return token.SignedString([]byte(SignKey))
}

func (us *UserServiceImp) GenerateGoogleToken(ctx context.Context, id int) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "stonks",
		},
		id,
		"google",
	})
	return token.SignedString([]byte(SignKey))
}

func (us *UserServiceImp) ParseToken(token string) (int, error) {
	return 0, nil
}
