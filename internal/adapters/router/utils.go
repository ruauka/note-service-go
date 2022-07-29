package router

import (
	"encoding/json"
	"log"
	"net/http"
)

//type respError struct {
//	Message string `json:"message"`
//}

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

