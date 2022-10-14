// Package middleware Package middleware
package middleware

import (
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"

	"web/internal/domain/errors"
	"web/internal/domain/services"
	"web/internal/utils/dictionary"
	"web/internal/utils/functions"
)

// CheckToken - handler middleware. Check bearer token for auth.
func CheckToken(next httprouter.Handle, auth services.UserAuthService) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		header := r.Header.Get("Authorization")
		if header == "" {
			functions.Abort(r.Context(), w, http.StatusUnauthorized, nil, errors.ErrEmptyAuthHeader, "", "")
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != dictionary.LenHeaderParts {
			functions.Abort(r.Context(), w, http.StatusUnauthorized, nil, errors.ErrInvalidAuthHeader, "", "")
			return
		}

		if headerParts[0] != "Bearer" {
			functions.Abort(r.Context(), w, http.StatusUnauthorized, nil, errors.ErrInvalidAuthHeader, "", "")
			return
		}

		if headerParts[1] == "" {
			functions.Abort(r.Context(), w, http.StatusUnauthorized, nil, errors.ErrEmptyToken, "", "")
			return
		}

		userID, err := auth.ParseToken(headerParts[1])
		if err != nil {
			functions.Abort(r.Context(), w, http.StatusUnauthorized, nil, err, "", "")
			return
		}

		r.Header.Set("user_id", userID)

		next(w, r, ps)
	}
}
