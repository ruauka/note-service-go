package services

import (
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt"

	"web/internal/adapters/storage"
	"web/internal/domain/enteties/dto"
	"web/internal/domain/enteties/model"
	"web/internal/utils"
)

type authService struct {
	storage storage.UserAuthStorage
	// logger
}

func NewAuthService(userAuthStorage storage.UserAuthStorage) UserAuthService {
	return &authService{storage: userAuthStorage}
}

func (a *authService) RegisterUser(user *model.User) (*model.User, error) {
	user.Password = utils.GeneratePasswordHash(user.Password)
	return a.storage.RegisterUser(user)
}

func (a *authService) GenerateToken(userName, password string) (string, error) {
	user, err := a.storage.GetUserForToken(userName, utils.GeneratePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, dto.TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(utils.ExpDuration).Unix(),
		},
		UserID: user.ID,
	})

	tokenString, err := token.SignedString([]byte(utils.SigningKey))
	if err != nil {
		log.Println(err)
	}

	return tokenString, nil
}

func (a *authService) ParseToken(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &dto.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(utils.SigningKey), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*dto.TokenClaims)
	if !ok {
		return "", errors.New("token claims are not of type *dto.TokenClaims")
	}

	return claims.UserID, nil
}
