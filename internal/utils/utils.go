// Package utils Package utils
package utils

import (
	"context"
	"crypto/sha1" //nolint:gosec
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/julienschmidt/httprouter"

	"web/internal/domain/errors"
	"web/pkg/logger"
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
	salt        = "abc"
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

// GeneratePasswordHash user password hash generator.
func GeneratePasswordHash(password string) string {
	hash := sha1.New() //nolint:gosec
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

// MakeJSONResponse make http response to UI.
func MakeJSONResponse(w http.ResponseWriter, httpStatus int, resp interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode(resp) //nolint:errcheck,gosec
}

// Abort func.
func Abort(ctx context.Context, w http.ResponseWriter, httpStatus int, err, errDesc error, name, instance string) {
	if err != nil {
		if instanceErr := ErrDBCheck(err.Error(), name, instance); instanceErr != nil {
			ErrCheck(ctx, w, httpStatus, nil, instanceErr)
			return
		}
		ErrCheck(ctx, w, httpStatus, err, errDesc)
		return
	}
	ErrCheck(ctx, w, httpStatus, err, errDesc)
}

// ErrCheck - make resp to UI.
func ErrCheck(ctx context.Context, w http.ResponseWriter, httpStatus int, err, errDesc error) {
	//nolint:errcheck,gosec
	json.NewEncoder(SetErrRespHeaders(w, httpStatus)).Encode(MapErrCreate(ctx, err, errDesc))
}

// SetErrRespHeaders set err in headers.
func SetErrRespHeaders(w http.ResponseWriter, httpStatus int) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	return w
}

// MapErrCreate create map with err.
func MapErrCreate(ctx context.Context, err, errDesc error) map[string]string {
	errMap := make(map[string]string)
	if err == nil {
		errMap["error"] = errDesc.Error()
		logger.LogFromContext(ctx).Error(errDesc.Error())
	} else {
		errMap["error"] = errDesc.Error()
		errMap["desc"] = err.Error()
		logger.LogFromContext(ctx).Error(errDesc.Error())
	}
	return errMap
}

// ErrDBCheck check BD err.
func ErrDBCheck(dbErr, name, instance string) error {
	switch {
	case strings.Contains(dbErr, errors.ErrDBDuplicate):
		return fmt.Errorf(fmt.Sprintf("%s '%s' is already exists", name, instance))
	case strings.Contains(dbErr, errors.ErrDBNotExists) && checkID(instance):
		return fmt.Errorf(fmt.Sprintf("No %s with id '%s'", name, instance))
	case strings.Contains(dbErr, errors.ErrDBNotExists):
		return fmt.Errorf(fmt.Sprintf("No %s with name '%s'", name, instance))
	}

	return nil
}

// checkID check try to int conversion.
func checkID(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}
