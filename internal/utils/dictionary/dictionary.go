// Package dictionary Package dictionary
package dictionary

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/julienschmidt/httprouter"
)

// users URLs.
const (
	UsersURL = "/users"
	UserURL  = "/users/:id"
	Register = "/register"
	Login    = "/login"
)

// notes URLs.
const (
	NotesURL       = "/notes"
	NoteURL        = "/notes/:id"
	AllTagsByNotes = "/allnotes/tags"
	AllTagsByNote  = "/notes/:id/tags"
)

// tags URLs.
const (
	TagsURL = "/tags"
	TagURL  = "/tags/:id"
)

// TagSet URLs.
const (
	TagsSet    = "/notes/:id/tags/set"
	TagsRemove = "/notes/:id/tags/remove"
)

// secret info.
const (
	Salt        = "abc"
	SigningKey  = "secret"
	ExpDuration = time.Hour * 12
)

// tables names for queries in storage.
const (
	UsersTable     = "users"
	NotesTable     = "notes"
	TagsTable      = "tags"
	NotesTagsTable = "notes_tags"
)

// for err.
const (
	User = "user"
	Note = "note"
	Tag  = "tag"
)

// LenHeaderParts len of arr.
const (
	LenHeaderParts = 2
)

// TokenClaims - additional token field to userID.
type TokenClaims struct {
	jwt.StandardClaims
	UserID string `json:"user_id"`
}

// Validate validate obj.
var Validate = validator.New()

// LogMiddleware custom type of logging middleware.
type LogMiddleware func(next httprouter.Handle) httprouter.Handle
