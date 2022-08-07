package errors

import "errors"

var (
	ErrDbResponse  = errors.New("db response error")
	ErrDbDuplicate = "duplicate key value violates unique constraint"
	ErrDbNotExists = "no rows in result set"
)

var (
	ErrUsersListEmpty    = errors.New("no users")
	ErrEmptyAuthHeader   = errors.New("empty auth header")
	ErrEmptyToken        = errors.New("empty token")
	ErrInvalidAuthHeader = errors.New("invalid auth header")
	ErrSigningMethod     = errors.New("invalid signing method")
	ErrClaimsType        = errors.New("token claims are not of type *dto.TokenClaims")
)

var (
	ErrNotesListEmpty         = errors.New("no notes")
	ErrNotesListWithTagsEmpty = errors.New("no notes with tags")
)

var (
	ErrTagsListEmpty = errors.New("no tags")
	ErrTagNotExists  = errors.New("tag not exists")
)
