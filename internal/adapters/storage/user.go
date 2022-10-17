package storage

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"

	"web/internal/domain/entities/dto"
	"web/internal/utils/dictionary"
	"web/internal/utils/functions"
)

// userStorage user storage struct.
type userStorage struct {
	db *sqlx.DB
}

// NewUserStorage user storage func builder.
func NewUserStorage(db *sqlx.DB) UserStorage {
	return &userStorage{db: db}
}

// GetUserByID get user by id from DB.
func (u *userStorage) GetUserByID(id string) (*dto.UserResp, error) {
	var user dto.UserResp

	query := fmt.Sprintf("SELECT id, username FROM %s WHERE id=$1", dictionary.UsersTable)
	if err := u.db.Get(&user, query, id); err != nil {
		return nil, err
	}

	return &user, nil
}

// GetAllUsers get all users from DB.
func (u *userStorage) GetAllUsers() ([]dto.UserResp, error) {
	var users []dto.UserResp

	query := fmt.Sprintf("SELECT id, username FROM %s", dictionary.UsersTable)
	if err := u.db.Select(&users, query); err != nil {
		return nil, err
	}

	return users, nil
}

// UpdateUser update user by id in DB.
func (u *userStorage) UpdateUser(newUser *dto.UserUpdate, userID string) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 1

	if newUser.Username != nil {
		setValues = append(setValues, fmt.Sprintf("username=$%d", argID))
		args = append(args, *newUser.Username)
		argID++
	}

	if newUser.Password != nil {
		setValues = append(setValues, fmt.Sprintf("password=$%d", argID))
		args = append(args, functions.GeneratePasswordHash(*newUser.Password))
		argID++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", dictionary.UsersTable, setQuery, argID)
	args = append(args, userID)

	_, err := u.db.Exec(query, args...)

	return err
}

// DeleteUser delete user by id from DB.
func (u *userStorage) DeleteUser(userID string) (int, error) {
	var id int

	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1 RETURNING id", dictionary.UsersTable)
	if err := u.db.QueryRow(query, userID).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
