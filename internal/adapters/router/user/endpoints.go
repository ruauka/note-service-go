package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"web/internal/adapters/router"
	"web/internal/domain/enteties/dto"
	"web/internal/domain/enteties/model"
	"web/internal/domain/errors"
)

func (h *handler) RegisterUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	user := &model.User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	user, err := h.service.Auth.RegisterUser(user)
	if err != nil {
		router.Abort(w, http.StatusBadRequest, err, errors.ErrDbResponse)
		return
	}

	resp := make(map[string]string)
	resp[fmt.Sprintf("Создан пользователь '%s' с id", user.Username)] = user.ID

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func (h *handler) GenerateToken(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	user := dto.UserAuth{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	token, err := h.service.Auth.GenerateToken(user.Username, user.Password)
	if err != nil {
		router.Abort(w, http.StatusBadRequest, err, errors.ErrDbResponse)
		return
	}

	resp := make(map[string]string)
	resp["token"] = token

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func (h *handler) GetAllUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	users, err := h.service.User.GetAllUsers()
	if err != nil {
		router.Abort(w, http.StatusBadRequest, err, errors.ErrDbResponse)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func (h *handler) GetUserByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	user, err := h.service.User.GetUserByID(ps.ByName("id"))
	if err != nil {
		router.Abort(w, http.StatusBadRequest, err, errors.ErrDbResponse)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	newUser := &dto.UserUpdate{}
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	_, err := h.service.User.GetUserByID(ps.ByName("id"))
	if err != nil {
		router.Abort(w, http.StatusBadRequest, err, errors.ErrUserNotExists)
		return
	}

	err = h.service.User.UpdateUser(newUser, ps.ByName("id"))
	if err != nil {
		router.Abort(w, http.StatusBadRequest, err, errors.ErrDbResponse)
		return
	}

	resp := make(map[string]string)
	resp["Обновлен пользователь с id"] = ps.ByName("id")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := h.service.User.DeleteUser(ps.ByName("id"))
	if err != nil {
		router.Abort(w, http.StatusBadRequest, err, errors.ErrDbResponse)
		return
	}

	resp := make(map[string]int)
	resp["Удален пользователь с id"] = id

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
