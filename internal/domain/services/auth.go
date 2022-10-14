package services

import (
	"time"

	"github.com/golang-jwt/jwt"

	"web/internal/adapters/storage"
	"web/internal/domain/entities/model"
	"web/internal/domain/errors"
	"web/internal/utils/dictionary"
	"web/internal/utils/functions"
)

// authService auth service struct.
type authService struct {
	storage storage.UserAuthStorage
}

// NewAuthService auth service func builder.
func NewAuthService(userAuthStorage storage.UserAuthStorage) UserAuthService {
	return &authService{storage: userAuthStorage}
}

// RegisterUser create user.
func (a *authService) RegisterUser(user *model.User) (*model.User, error) {
	user.Password = functions.GeneratePasswordHash(user.Password)
	return a.storage.RegisterUser(user)
}

// GenerateToken generate token for user auth.
func (a *authService) GenerateToken(userName, password string) (string, error) {
	user, err := a.storage.GetUserForToken(userName, functions.GeneratePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, dictionary.TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(dictionary.ExpDuration).Unix(),
		},
		UserID: user.ID,
	})

	tokenString, err := token.SignedString([]byte(dictionary.SigningKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseToken check token for auth.
func (a *authService) ParseToken(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &dictionary.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.ErrSigningMethod
		}
		return []byte(dictionary.SigningKey), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*dictionary.TokenClaims)
	if !ok {
		return "", errors.ErrClaimsType
	}

	return claims.UserID, nil
}
