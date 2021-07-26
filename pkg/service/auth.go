package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/DimKush/guestbook/tree/main/internal/entities/User"
	"github.com/DimKush/guestbook/tree/main/pkg/repository"
)

const salt = "ssgsdgdfggegrgwgwefwefwefwefdf4r231"

type AuthService struct {
	repos repository.Authorization
}

func InitAuthService(repos repository.Authorization) *AuthService {
	return &AuthService{repos: repos}
}

func (data *AuthService) CreateUser(user User.User) (int, error) {
	user.Password = data.generatePassHash(user.Password)
	return data.repos.CreateUser(user)
}

func (data *AuthService) generatePassHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
