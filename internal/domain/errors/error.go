package errors

import "errors"

var (
	ErrDbResponse = errors.New("db response error")
	ErrUserNotExists = errors.New("user not exists")
	ErrUsersListEmpty = errors.New("no users")
)
