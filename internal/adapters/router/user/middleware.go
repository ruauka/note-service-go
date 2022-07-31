package user

import (
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"

	"web/internal/domain/errors"
	"web/internal/utils"
)

func (h *handler) CheckToken(fn httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		header := r.Header.Get("Authorization")
		if header == "" {
			utils.Abort(w, http.StatusUnauthorized, nil, errors.ErrEmptyAuthHeader)
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 {
			utils.Abort(w, http.StatusUnauthorized, nil, errors.ErrInvalidAuthHeader)
			return
		}

		userId, err := h.service.Auth.ParseToken(headerParts[1])
		if err != nil {
			utils.Abort(w, http.StatusUnauthorized, nil, err)
			return
		}

		r.Header.Set("user_id", userId)

		fn(w, r, ps)
	}
}
