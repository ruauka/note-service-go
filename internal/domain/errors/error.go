// Package errors Package errors
package errors

import "errors"

// DB errors.
var (
	ErrDBResponse  = errors.New("db response error")
	ErrDBDuplicate = "duplicate key value violates unique constraint"
	ErrDBNotExists = "no rows in result set"
)

// user errors.
var (
	ErrUsersListEmpty    = errors.New("no users")
	ErrEmptyAuthHeader   = errors.New("empty auth header")
	ErrEmptyToken        = errors.New("empty token")
	ErrInvalidAuthHeader = errors.New("invalid auth header")
	ErrSigningMethod     = errors.New("invalid signing method")
	ErrClaimsType        = errors.New("token claims are not of type *dto.TokenClaims")
)

// notes errors.
var (
	ErrNotesListEmpty         = errors.New("no notes")
	ErrNotesListWithTagsEmpty = errors.New("no notes with tags")
)

// ErrTagsListEmpty tags errors.
var ErrTagsListEmpty = errors.New("no tags")
