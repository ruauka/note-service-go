package services

import (
	"web/internal/adapters/storage"
	"web/internal/domain/interfaces"
	"web/internal/domain/services/user"
)

type Services struct {
	Auth interfaces.UserAuthService
	User interfaces.UserService
}

func NewServices(db *storage.Storages) *Services {
	return &Services{
		Auth: user.NewAuthService(db.Auth),
		User: user.NewUserService(db.User),
	}
}
