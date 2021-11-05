package hash

import (
	"context"
	"github.com/veremchukvv/stonks-test/pkg/logging"
	"golang.org/x/crypto/bcrypt"
)

type PasswordHasher interface {
	Hash(password string) (string, error)
	CheckPWD(password string, hash string) (bool, error)
}

type BCHasher struct {
	ctx context.Context
}

func NewBCPasswordHasher(ctx context.Context) *BCHasher {
	return &BCHasher{
		ctx,
	}
}

func (bch *BCHasher) Hash(password string) (string, error) {
	log := logging.FromContext(bch.ctx)
	pwd := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(pwd, 14)
	if err != nil {
		log.Error("Can't generate password hash")
		return "", err
	}
	return string(hash), nil
}

func (bch *BCHasher) CheckPWD(password string, hash string) (bool, error) {
	log := logging.FromContext(bch.ctx)
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		log.Infof("Password incorrect %v", err)
		return false, err
	}
	return true, nil
}
