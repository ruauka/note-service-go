package services

import (
	"web/internal/domain/interfaces"
	"web/internal/domain/services/user"
)

type Services struct {
	Auth interfaces.UserAuthService
	User interfaces.UserService
}

func NewServices(db interfaces.Storage) *Services {
	return &Services{
		Auth: user.NewAuthService(db),
		User: user.NewUserService(db),
	}
}
