package utils

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	UsersURL = "/users"
	UserURL  = "/users/:id"
	Register = "/register"
	Auth     = "/auth"
)

const (
	NotesURL = "/notes"
	NoteURL  = "/notes/:id"
)

const (
	TagsURL = "/tags"
	TagURL  = "/tags/:id"
)

const (
	SigningKey  = "secret"
	ExpDuration = time.Hour * 12
)

const (
	UsersTable = "users"
	NotesTable = "notes"
	TagsTable  = "tags"
)

const salt = "abc"

func GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

// SetErrRespHeaders - установка необходимых хедеров для ответа с ошибкой.
func SetErrRespHeaders(w http.ResponseWriter, httpStatus int) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	return w
}

// MapErrCreate - Создание словаря с ошибкой для ответа UI.
func MapErrCreate(err, errDesc error) map[string]string {
	errMap := make(map[string]string)
	if err == nil {
		errMap["error"] = errDesc.Error()
	} else {
		errMap["error"] = errDesc.Error()
		errMap["desc"] = err.Error()
	}
	return errMap
}

// Abort - ответ UI.
func Abort(w http.ResponseWriter, httpStatus int, err, errDesc error) {
	// nolint:errcheck,gosec
	json.NewEncoder(SetErrRespHeaders(w, httpStatus)).Encode(MapErrCreate(err, errDesc))
	log.Println(errDesc.Error())
}
