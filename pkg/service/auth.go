package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"net/mail"
	"regexp"
	"sync"
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
	//regexpUsername = ""
	//regexpEmail =	 "^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$"
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

func (data *AuthService) checkFilledUser(user *User.User) error {
	fmt.Println("HERE!")
	var wg sync.WaitGroup
	regexpChan := make(chan error, 2)

	wg.Add(2)
	go func(email string) {
		defer wg.Done()
		if _, err := mail.ParseAddress(email); err != nil {
			regexpChan <- fmt.Errorf("Email %s is invalid.", email)
		}
	}(user.Email)

	go func(username string) {
		defer wg.Done()
		usernameRegexp := regexp.MustCompile(`^[a-zA-Z0-9_]{5,}[a-zA-Z]+[0-9]*$`)
		if !usernameRegexp.MatchString(username) {
			regexpChan <- fmt.Errorf("Username %s is invalid.", username)
		}
	}(user.Username)

	wg.Wait()
	close(regexpChan)

	var res string
	for errVal := range regexpChan {
		if errVal != nil {
			res += errVal.Error()
			res += "\n"
		}
	}

	fmt.Println(res)

	if res != "" {
		return fmt.Errorf(res)
	} else {
		return nil
	}
}

func (data *AuthService) CreateUser(user User.User) (int, error) {
	if err := data.checkFilledUser(&user); err != nil {
		return 0, err
	}

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

func (data *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Invalid string methods.")
		}

		return []byte(signingKey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("Error during parse the token. Token claims are not of types.")
	}

	return claims.UserId, nil
}
