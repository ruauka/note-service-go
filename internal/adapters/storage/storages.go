package storage

import (
	"github.com/jmoiron/sqlx"

	user2 "web/internal/adapters/storage/user"
	"web/internal/domain/interfaces"
)

type Storages struct {
	Auth interfaces.UserAuthStorage
	User interfaces.UserStorage
}

func NewStorages(pgDB *sqlx.DB) *Storages {
	return &Storages{
		Auth: user2.NewAuthStorage(pgDB),
		User: user2.NewUserStorage(pgDB),
	}
}
