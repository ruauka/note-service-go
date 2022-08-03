package middleware

import (
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"

	"web/internal/domain/errors"
	"web/internal/domain/services"
	"web/internal/utils"
)

func CheckToken(next httprouter.Handle, auth services.UserAuthService) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		header := r.Header.Get("Authorization")
		if header == "" {
			utils.ErrCheck(w, http.StatusUnauthorized, nil, errors.ErrEmptyAuthHeader)
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 {
			utils.ErrCheck(w, http.StatusUnauthorized, nil, errors.ErrInvalidAuthHeader)
			return
		}

		userId, err := auth.ParseToken(headerParts[1])
		if err != nil {
			utils.ErrCheck(w, http.StatusUnauthorized, nil, err)
			return
		}

		r.Header.Set("user_id", userId)

		next(w, r, ps)
	}
}
