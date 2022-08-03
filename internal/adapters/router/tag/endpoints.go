package tag

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

func (h *handler) CreateTag(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tag := &model.Tag{}
	if err := json.NewDecoder(r.Body).Decode(&tag); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	userID := r.Header.Get("user_id")

	tag, err := h.service.Tag.CreateTag(tag, userID)
	if err != nil {
		utils.ErrCheck(w, http.StatusBadRequest, err, errors.ErrDbResponse)
		return
	}

	resp := make(map[string]string)
	resp[fmt.Sprintf("Создан тэг '%s' с id", tag.TagName)] = tag.ID

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func (h *handler) GetTagByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID := r.Header.Get("user_id")

	tag, err := h.service.Tag.GetTagByID(ps.ByName("id"), userID)
	if err != nil {
		utils.ErrCheck(w, http.StatusBadRequest, err, errors.ErrDbResponse)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tag)
}

func (h *handler) GetAllTagsByUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	userID := r.Header.Get("user_id")

	tags, err := h.service.Tag.GetAllTagsByUser(userID)
	if err != nil {
		utils.ErrCheck(w, http.StatusBadRequest, err, errors.ErrDbResponse)
		return
	}

	if len(tags) == 0 {
		utils.ErrCheck(w, http.StatusBadRequest, err, errors.ErrTagsListEmpty)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tags)
}

func (h *handler) UpdateTag(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	tag := &dto.TagUpdate{}
	if err := json.NewDecoder(r.Body).Decode(&tag); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	userID := r.Header.Get("user_id")
	tagID := ps.ByName("id")

	_, err := h.service.Tag.GetTagByID(tagID, userID)
	if err != nil {
		utils.ErrCheck(w, http.StatusBadRequest, err, errors.ErrTagNotExists)
		return
	}

	err = h.service.Tag.UpdateTag(tag, tagID)
	if err != nil {
		utils.ErrCheck(w, http.StatusBadRequest, err, errors.ErrDbResponse)
		return
	}

	resp := make(map[string]string)
	resp["Обновлен тэг с id"] = tagID

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *handler) DeleteTag(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID := r.Header.Get("user_id")

	tagID, err := h.service.Tag.DeleteTag(ps.ByName("id"), userID)
	if err != nil {
		utils.ErrCheck(w, http.StatusBadRequest, err, errors.ErrDbResponse)
		return
	}

	resp := make(map[string]int)
	resp["Удален тэг с id"] = tagID

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
