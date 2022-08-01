package errors

import "errors"

var (
	ErrDbResponse        = errors.New("db response error")
	ErrUserNotExists     = errors.New("user not exists")
	ErrUsersListEmpty    = errors.New("no users")
	ErrEmptyAuthHeader   = errors.New("empty auth header")
	ErrInvalidAuthHeader = errors.New("invalid auth header")
	ErrNotesListEmpty    = errors.New("no users")
	ErrNoteNotExists     = errors.New("note not exists")
)
