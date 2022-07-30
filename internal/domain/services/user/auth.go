package user

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt"

	"web/internal/config"
	"web/internal/domain/enteties/dto"
	"web/internal/domain/enteties/model"
	"web/internal/domain/interfaces"
	"web/internal/utils"
)

type authService struct {
	storage interfaces.UserAuthStorage
	// logger
}

func NewAuthService(db interfaces.UserAuthStorage) interfaces.UserAuthService {
	return &authService{storage: db}
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
			ExpiresAt: time.Now().Add(config.ExpDuration).Unix(),
		},
		UserID: user.ID,
	})

	tokenString, err := token.SignedString([]byte(config.SigningKey))
	if err != nil {
		log.Println(err)
	}

	return tokenString, nil
}
