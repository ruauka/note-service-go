package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"web/internal/adapters/router/validate"
	"web/internal/domain/entities/dto"
	"web/internal/domain/entities/model"
	"web/internal/domain/errors"
	"web/internal/utils/dictionary"
	"web/internal/utils/functions"
	"web/pkg/logger"
)

// RegisterUser create user.
func (h *Handler) RegisterUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := r.Context()

	newUser := &model.User{}
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}
	// Валидация объекта структуры User //
	err := validate.InputJSONValidate(newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	user, err := h.service.Auth.RegisterUser(newUser)
	if err != nil {
		functions.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrDBResponse, dictionary.User, newUser.Username)
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	resp := make(map[string]string)
	resp[fmt.Sprintf("Created new user '%s' with id", user.Username)] = user.ID

	functions.MakeJSONResponse(w, http.StatusCreated, resp)
}

// GenerateToken generate token for user auth.
func (h *Handler) GenerateToken(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := r.Context()

	user := dto.UserAuth{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}
	// Валидация объекта структуры UserAuth //
	err := validate.InputJSONValidate(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	token, err := h.service.Auth.GenerateToken(user.Username, user.Password)
	if err != nil {
		functions.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrDBResponse, dictionary.User, user.Username)
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	resp := make(map[string]string)
	resp["token"] = fmt.Sprintf("Bearer %s", token)

	functions.MakeJSONResponse(w, http.StatusOK, resp)
}

// GetUserByID get user by ID.
func (h *Handler) GetUserByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID := ps.ByName("id")
	ctx := r.Context()

	user, err := h.service.User.GetUserByID(userID)
	if err != nil {
		functions.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrDBResponse, dictionary.User, userID)
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	functions.MakeJSONResponse(w, http.StatusOK, user)
}

// GetAllUsers get all users.
func (h *Handler) GetAllUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := r.Context()

	users, err := h.service.User.GetAllUsers()
	if err != nil {
		functions.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrDBResponse, "", "")
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	if len(users) == 0 {
		functions.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrUsersListEmpty, "", "")
		logger.LogFromContext(ctx).Error(errors.ErrUsersListEmpty.Error())
		return
	}

	functions.MakeJSONResponse(w, http.StatusOK, users)
}

// UpdateUser update user by ID.
func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID := ps.ByName("id")
	ctx := r.Context()

	newUser := &dto.UserUpdate{}
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	_, err := h.service.User.GetUserByID(userID)
	if err != nil {
		functions.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrDBResponse, dictionary.User, userID)
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	err = h.service.User.UpdateUser(newUser, userID)
	if err != nil {
		functions.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrDBResponse, dictionary.User, *newUser.Username)
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	resp := make(map[string]string)
	resp["Updated user with id"] = userID

	functions.MakeJSONResponse(w, http.StatusOK, resp)
}

// DeleteUser delete user by ID.
func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID := ps.ByName("id")
	ctx := r.Context()

	_, err := h.service.User.GetUserByID(userID)
	if err != nil {
		functions.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrDBResponse, dictionary.User, userID)
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	id, err := h.service.User.DeleteUser(userID)
	if err != nil {
		functions.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrDBResponse, "", "")
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	resp := make(map[string]int)
	resp["Deleted user with id"] = id

	functions.MakeJSONResponse(w, http.StatusOK, resp)
}
