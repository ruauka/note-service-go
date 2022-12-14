package storage

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"web/internal/domain/entities/model"
	"web/internal/utils/dictionary"
)

// userAuthStorage auth storage struct.
type userAuthStorage struct {
	db *sqlx.DB
}

// NewAuthStorage auth storage func builder.
func NewAuthStorage(db *sqlx.DB) UserAuthStorage {
	return &userAuthStorage{db: db}
}

// RegisterUser insert user in DB.
func (s *userAuthStorage) RegisterUser(user *model.User) (*model.User, error) {
	query := fmt.Sprintf("INSERT INTO %s (username, password) VALUES ($1, $2) RETURNING id", dictionary.UsersTable)
	if err := s.db.QueryRow(query, user.Username, user.Password).Scan(&user.ID); err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserForToken get user from DB for token gen func.
func (s *userAuthStorage) GetUserForToken(userName, passwordHash string) (*model.User, error) {
	var user model.User

	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password=$2", dictionary.UsersTable)
	err := s.db.Get(&user, query, userName, passwordHash)

	return &user, err
}
