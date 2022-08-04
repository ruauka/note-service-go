package utils

import (
	"context"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"

	"web/internal/domain/errors"
	"web/pkg/logger"
)

const (
	UsersURL = "/users"
	UserURL  = "/users/:id"
	Register = "/register"
	Login    = "/login"
)

const (
	NotesURL       = "/notes"
	NoteURL        = "/notes/:id"
	AllTagsByNotes = "/allnotes/tags"
	AllTagsByNote  = "/notes/:id/tags"
)

const (
	TagsURL = "/tags"
	TagURL  = "/tags/:id"
)

const (
	TagsSet    = "/notes/:id/tags/set"
	TagsRemove = "/notes/:id/tags/remove"
)

const (
	SigningKey  = "secret"
	ExpDuration = time.Hour * 12
)

const (
	UsersTable     = "users"
	NotesTable     = "notes"
	TagsTable      = "tags"
	NotesTagsTable = "notes_tags"
)

const (
	User = "user"
	Note = "note"
	Tag  = "tag"
)

const salt = "abc"

var Validate = validator.New()

func GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func MakeJsonResponse(w http.ResponseWriter, httpStatus int, resp interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode(resp)
}

func Abort(ctx context.Context, w http.ResponseWriter, httpStatus int, err, errDesc error, name, instance string) {
	if err != nil {
		if instanceErr := ErrDbCheck(err.Error(), name, instance); instanceErr != nil {
			ErrCheck(ctx, w, httpStatus, nil, instanceErr)
			return
		} else {
			ErrCheck(ctx, w, httpStatus, err, errDesc)
			return
		}
	} else {
		ErrCheck(ctx, w, httpStatus, err, errDesc)
		return
	}
}

// ErrCheck - ответ UI.
func ErrCheck(ctx context.Context, w http.ResponseWriter, httpStatus int, err, errDesc error) {
	// nolint:errcheck,gosec
	json.NewEncoder(SetErrRespHeaders(w, httpStatus)).Encode(MapErrCreate(ctx, err, errDesc))
}

// SetErrRespHeaders - установка необходимых хедеров для ответа с ошибкой.
func SetErrRespHeaders(w http.ResponseWriter, httpStatus int) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	return w
}

// MapErrCreate - Создание словаря с ошибкой для ответа UI.
func MapErrCreate(ctx context.Context, err, errDesc error) map[string]string {
	errMap := make(map[string]string)
	if err == nil {
		errMap["error"] = errDesc.Error()
		logger.LogFromContext(ctx).Error(errDesc.Error())
	} else {
		errMap["error"] = errDesc.Error()
		errMap["desc"] = err.Error()
		logger.LogFromContext(ctx).Error(err.Error())
		logger.LogFromContext(ctx).Error(errDesc.Error())
	}
	return errMap
}

func ErrDbCheck(dbErr, name, instance string) error {
	switch {
	case strings.Contains(dbErr, errors.ErrDbDuplicate):
		return fmt.Errorf(fmt.Sprintf("%s '%s' is already exists", name, instance))
	case strings.Contains(dbErr, errors.ErrDbNotExists) && checkID(instance):
		return fmt.Errorf(fmt.Sprintf("No %s with id '%s'", name, instance))
	case strings.Contains(dbErr, errors.ErrDbNotExists):
		return fmt.Errorf(fmt.Sprintf("No %s with name '%s'", name, instance))
	}

	return nil
}

func checkID(str string) bool {
	_, err := strconv.Atoi(str)
	if err != nil {
		return false
	}
	return true
}
