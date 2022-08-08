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
			utils.Abort(r.Context(), w, http.StatusUnauthorized, nil, errors.ErrEmptyAuthHeader, "", "")
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 {
			utils.Abort(r.Context(), w, http.StatusUnauthorized, nil, errors.ErrInvalidAuthHeader, "", "")
			return
		}

		if headerParts[0] != "Bearer" {
			utils.Abort(r.Context(), w, http.StatusUnauthorized, nil, errors.ErrInvalidAuthHeader, "", "")
			return
		}

		if headerParts[1] == "" {
			utils.Abort(r.Context(), w, http.StatusUnauthorized, nil, errors.ErrEmptyToken, "", "")
			return
		}

		userId, err := auth.ParseToken(headerParts[1])
		if err != nil {
			utils.Abort(r.Context(), w, http.StatusUnauthorized, nil, err, "", "")
			return
		}

		r.Header.Set("user_id", userId)

		next(w, r, ps)
	}
}
