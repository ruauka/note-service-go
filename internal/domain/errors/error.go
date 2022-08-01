package errors

import "errors"

var (
	ErrDbResponse = errors.New("db response error")
)

var (
	ErrUsersListEmpty    = errors.New("no users")
	ErrUserNotExists     = errors.New("user not exists")
	ErrEmptyAuthHeader   = errors.New("empty auth header")
	ErrInvalidAuthHeader = errors.New("invalid auth header")
)

var (
	ErrNotesListEmpty = errors.New("no notes")
	ErrNoteNotExists  = errors.New("note not exists")
)
