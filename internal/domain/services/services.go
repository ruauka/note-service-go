package services

import (
	"web/internal/domain/interfaces"
	"web/internal/domain/services/user"
)

type Services struct {
	User interfaces.UserService
}

func NewServices(db interfaces.Storage) *Services {
	return &Services{
		User: user.NewUserService(db),
	}
}
