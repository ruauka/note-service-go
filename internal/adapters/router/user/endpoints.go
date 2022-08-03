package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"web/internal/domain/enteties/dto"
	"web/internal/domain/enteties/model"
	"web/internal/domain/errors"
	"web/internal/utils"
)

func (h *handler) RegisterUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	newUser := &model.User{}
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	user, err := h.service.Auth.RegisterUser(newUser)
	if err != nil {
		if checkErr := utils.CheckDbErr(err.Error(), utils.User, newUser.Username); checkErr != nil {
			utils.Abort(w, http.StatusBadRequest, nil, checkErr)
			return
		} else {
			utils.Abort(w, http.StatusBadRequest, err, errors.ErrDbResponse)
			return
		}
	}

	resp := make(map[string]string)
	resp[fmt.Sprintf("Создан пользователь '%s' с id", user.Username)] = user.ID

	utils.MakeJsonResponse(w, http.StatusCreated, resp)
}

func (h *handler) GenerateToken(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	user := dto.UserAuth{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	token, err := h.service.Auth.GenerateToken(user.Username, user.Password)
	if err != nil {
		if checkErr := utils.CheckDbErr(err.Error(), utils.User, user.Username); checkErr != nil {
			utils.Abort(w, http.StatusBadRequest, nil, checkErr)
			return
		} else {
			utils.Abort(w, http.StatusBadRequest, err, errors.ErrDbResponse)
			return
		}
	}

	resp := make(map[string]string)
	resp["token"] = token

	utils.MakeJsonResponse(w, http.StatusOK, resp)
}

func (h *handler) GetAllUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	users, err := h.service.User.GetAllUsers()
	if err != nil {
		utils.Abort(w, http.StatusBadRequest, err, errors.ErrDbResponse)
		return
	}

	if len(users) == 0 {
		utils.Abort(w, http.StatusBadRequest, err, errors.ErrUsersListEmpty)
		return
	}

	utils.MakeJsonResponse(w, http.StatusOK, users)
}

func (h *handler) GetUserByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID := ps.ByName("id")

	user, err := h.service.User.GetUserByID(userID)
	if err != nil {
		if checkErr := utils.CheckDbErr(err.Error(), utils.User, userID); checkErr != nil {
			utils.Abort(w, http.StatusBadRequest, nil, checkErr)
			return
		} else {
			utils.Abort(w, http.StatusBadRequest, err, errors.ErrDbResponse)
			return
		}
	}

	utils.MakeJsonResponse(w, http.StatusOK, user)
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	newUser := &dto.UserUpdate{}
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	userID := ps.ByName("id")

	_, err := h.service.User.GetUserByID(userID)
	if err != nil {
		if checkErr := utils.CheckDbErr(err.Error(), utils.User, userID); checkErr != nil {
			utils.Abort(w, http.StatusBadRequest, nil, checkErr)
			return
		} else {
			utils.Abort(w, http.StatusBadRequest, err, errors.ErrDbResponse)
			return
		}
	}

	err = h.service.User.UpdateUser(newUser, userID)
	if err != nil {
		utils.Abort(w, http.StatusBadRequest, err, errors.ErrDbResponse)
		return
	}

	resp := make(map[string]string)
	resp["Обновлен пользователь с id"] = userID

	utils.MakeJsonResponse(w, http.StatusOK, resp)
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID := ps.ByName("id")

	_, err := h.service.User.GetUserByID(userID)
	if err != nil {
		if checkErr := utils.CheckDbErr(err.Error(), utils.User, userID); checkErr != nil {
			utils.Abort(w, http.StatusBadRequest, nil, checkErr)
			return
		} else {
			utils.Abort(w, http.StatusBadRequest, err, errors.ErrDbResponse)
			return
		}
	}

	id, err := h.service.User.DeleteUser(userID)
	if err != nil {
		utils.Abort(w, http.StatusBadRequest, err, errors.ErrDbResponse)
		return
	}

	resp := make(map[string]int)
	resp["Удален пользователь с id"] = id

	utils.MakeJsonResponse(w, http.StatusOK, resp)
}
