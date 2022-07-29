package storage

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"

	"web/internal/domain/enteties/dto"
	"web/internal/domain/enteties/model"
	"web/internal/domain/interfaces"
	"web/internal/domain/services/utils"
)

const (
	UsersTable = "users"
)

type storage struct {
	db *sqlx.DB
}

func NewStorage(pgDB *sqlx.DB) interfaces.Storage {
	return &storage{db: pgDB}
}

func (s *storage) RegisterUser(user *model.User) (*model.User, error) {
	query := fmt.Sprintf("INSERT INTO %s (username, password) VALUES ($1, $2) RETURNING id", UsersTable)
	if err := s.db.QueryRow(query, user.Username, user.Password).Scan(&user.ID); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *storage) GetUserForAuth(userName, password string) (*model.User, error) {
	var user model.User

	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password=$2", UsersTable)
	err := s.db.Get(&user, query, userName, password)

	return &user, err
}

func (s *storage) GetUserByID(id string) (*dto.UserResp, error) {
	var user dto.UserResp

	query := fmt.Sprintf("SELECT id, username FROM %s WHERE id=$1", UsersTable)
	if err := s.db.Get(&user, query, id); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *storage) GetAllUsers() ([]dto.UserResp, error) {
	var users []dto.UserResp

	getAllUsersQuery := fmt.Sprintf("SELECT id, username FROM %s", UsersTable)
	if err := s.db.Select(&users, getAllUsersQuery); err != nil {
		return nil, err
	}

	return users, nil
}

func (s *storage) UpdateUser(newUser *dto.UserUpdate, userId string) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if newUser.Username != nil {
		setValues = append(setValues, fmt.Sprintf("username=$%d", argId))
		args = append(args, *newUser.Username)
		argId++
	}

	if newUser.Password != nil {
		setValues = append(setValues, fmt.Sprintf("password=$%d", argId))
		args = append(args, utils.GeneratePasswordHash(*newUser.Password))
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", UsersTable, setQuery, argId)
	args = append(args, userId)

	_, err := s.db.Exec(query, args...)
	return err
}

func (s *storage) DeleteUser(userId string) (int, error) {
	var id int

	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1 RETURNING id", UsersTable)
	if err := s.db.QueryRow(query, userId).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
