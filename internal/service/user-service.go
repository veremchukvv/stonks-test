package service

import (
	"context"
	"github.com/golang-jwt/jwt"
	"github.com/veremchukvv/stonks-test/internal/models"
	"github.com/veremchukvv/stonks-test/internal/repository"
	"github.com/veremchukvv/stonks-test/pkg/hash"
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

var signKey = os.Getenv("SIGN_KEY")

type tokenClaims struct {
	jwt.StandardClaims
	UserId int
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
	})
	return token.SignedString([]byte(signKey))
}

func (us *UserServiceImp) GenerateVKToken(ctx context.Context, vkid int) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		vkid,
	})
	return token.SignedString([]byte(signKey))
}

func (us *UserServiceImp) ParseToken(token string) (int, error) {
	return 0, nil
}
