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

//type UserSignupReq struct {
//	Name     string
//	Email    string
//	Password string
//}
//
//type UserSigninReq struct {
//	Email    string
//	Password string
//}

const (
	tokenTTL = 12 * time.Hour
)

var SignKey = os.Getenv("SIGN_KEY")

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
	VKUserId int `json:"vkuser_id"`
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

func (us *UserServiceImp) GetUser(ctx context.Context, token string) (*models.User, *models.VKUser, error) {
	log := logging.FromContext(ctx)

	GetUserErr := errors.New("Can't get user")

	parsedToken, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(key *jwt.Token) (interface{}, error) {
		return []byte(SignKey), nil
		})
	if err != nil {
		return nil, nil, err
	}

	claims := parsedToken.Claims.(*tokenClaims)

	if claims.UserId != 0 {
		u, err := us.repo.GetUserByID(ctx, claims.UserId)
		if err != nil {
			log.Errorf("can't get user from db: %v", err)
			return nil, nil, err
		}
		return u, nil, nil
	}

	if claims.VKUserId != 0 {
		vku, err := us.repo.GetVKUserByID(ctx, claims.VKUserId)
		if err != nil {
			log.Errorf("can't get vkuser from db: %v", err)
			return nil, nil, err
		}
		return nil, vku, nil
	}

	return nil, nil, GetUserErr
}

func (us *UserServiceImp) GetVKUserByID(ctx context.Context, vkid int) (*models.VKUser, error) {
	return us.repo.GetVKUserByID(ctx, vkid)
}

func (us *UserServiceImp) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	var err error
	user.Password, err = us.hasher.Hash(user.Password)
	if err != nil {
		return nil, err
	}
	return us.repo.CreateUser(ctx, user)
}

func (us *UserServiceImp) CreateVKUser(ctx context.Context, vkuser *models.VKUser) (*models.VKUser, error) {
	return us.repo.CreateVKUser(ctx, vkuser)
}

func (us *UserServiceImp) GenerateToken(ctx context.Context, email string, password string) (string, error) {
	u, err := us.repo.GetUserByEmail(ctx, email)
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
		0,
	})
	return token.SignedString([]byte(SignKey))
}

func (us *UserServiceImp) GenerateVKToken(ctx context.Context, vkid int) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer: "stonks",
		},
		0,
		vkid,
	})
	return token.SignedString([]byte(SignKey))
}

func (us *UserServiceImp) ParseToken(token string) (int, error) {
	return 0, nil
}
