package storage

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"

	"web/internal/domain/enteties/dto"
	"web/internal/utils"
)

type userStorage struct {
	db *sqlx.DB
}

func NewUserStorage(pgDB *sqlx.DB) UserStorage {
	return &userStorage{db: pgDB}
}

func (u *userStorage) GetUserByID(id string) (*dto.UserResp, error) {
	var user dto.UserResp

	query := fmt.Sprintf("SELECT id, username FROM %s WHERE id=$1", utils.UsersTable)
	if err := u.db.Get(&user, query, id); err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userStorage) GetAllUsers() ([]dto.UserResp, error) {
	var users []dto.UserResp

	getAllUsersQuery := fmt.Sprintf("SELECT id, username FROM %s", utils.UsersTable)
	if err := u.db.Select(&users, getAllUsersQuery); err != nil {
		return nil, err
	}

	return users, nil
}

func (u *userStorage) UpdateUser(newUser *dto.UserUpdate, userId string) error {
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
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", utils.UsersTable, setQuery, argId)
	args = append(args, userId)

	_, err := u.db.Exec(query, args...)
	return err
}

func (u *userStorage) DeleteUser(userId string) (int, error) {
	var id int

	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1 RETURNING id", utils.UsersTable)
	if err := u.db.QueryRow(query, userId).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
