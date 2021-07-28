package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/DimKush/guestbook/tree/main/internal/entities/User"
	"github.com/DimKush/guestbook/tree/main/pkg/repository"
	"github.com/golang-jwt/jwt"
	"github.com/rs/zerolog/log"
)

const (
	salt       = "ssgsdgdfggegrgwgwefwefwefwefdf4r231"
	signingKey = "dzgjhhmnghty4T356cczxXzxcvxzbvvxcbgnfgnergeGWER"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}
type AuthService struct {
	auth repository.Authorization
}

func InitAuthService(repos repository.Authorization) *AuthService {
	return &AuthService{auth: repos}
}

func (data *AuthService) CreateUser(user User.User) (int, error) {
	user.Password = data.generatePassHash(user.Password)

	Audit.WriteEventParams("AuthService",
		"CreateUser",
		AUDIT_INFO,
		time.Now(),
		false,
		"Try to create user",
	)

	return data.auth.CreateUser(user)
}

func (data *AuthService) generatePassHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (data *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := data.auth.GetUser(username, password)
	if err != nil {
		log.Error().Msg(err.Error())
		return "", err
	}

	fmt.Println(user)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}
