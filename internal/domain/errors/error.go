package errors

import "errors"

var (
	ErrDbResponse  = errors.New("db response error")
	ErrDbDuplicate = "duplicate key value violates unique constraint"
	ErrDbNotExists = "no rows in result set"
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

var (
	ErrTagsListEmpty = errors.New("no tags")
	ErrTagNotExists  = errors.New("tag not exists")
)
