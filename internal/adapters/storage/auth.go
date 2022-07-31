package storage

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"web/internal/domain/enteties/model"
	"web/internal/utils"
)

type userAuthStorage struct {
	db *sqlx.DB
}

func NewAuthStorage(pgDB *sqlx.DB) UserAuthStorage {
	return &userAuthStorage{db: pgDB}
}

func (s *userAuthStorage) RegisterUser(user *model.User) (*model.User, error) {
	query := fmt.Sprintf("INSERT INTO %s (username, password) VALUES ($1, $2) RETURNING id", utils.UsersTable)
	if err := s.db.QueryRow(query, user.Username, user.Password).Scan(&user.ID); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userAuthStorage) GetUserForToken(userName, passwordHash string) (*model.User, error) {
	var user model.User

	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password=$2", utils.UsersTable)
	err := s.db.Get(&user, query, userName, passwordHash)

	return &user, err
}
