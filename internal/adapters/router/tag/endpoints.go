package tag

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"web/internal/adapters/router/validate"
	"web/internal/domain/enteties/dto"
	"web/internal/domain/enteties/model"
	"web/internal/domain/errors"
	"web/internal/utils"
)

func (h *handler) CreateTag(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	newTag := &model.Tag{}
	if err := json.NewDecoder(r.Body).Decode(&newTag); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	// Валидация объекта структуры Tag //
	err := validate.InputJsonValidate(newTag)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	userID := r.Header.Get("user_id")

	tag, err := h.service.Tag.CreateTag(newTag, userID)
	if err != nil {
		utils.Abort(nil, w, http.StatusBadRequest, err, errors.ErrDbResponse, utils.Tag, newTag.TagName)
		return
	}

	resp := make(map[string]string)
	resp[fmt.Sprintf("Создан тэг '%s' с id", tag.TagName)] = tag.ID

	utils.MakeJsonResponse(w, http.StatusCreated, resp)
}

func (h *handler) GetTagByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID := r.Header.Get("user_id")
	tagID := ps.ByName("id")

	tag, err := h.service.Tag.GetTagByID(tagID, userID)
	if err != nil {
		utils.Abort(nil, w, http.StatusBadRequest, err, errors.ErrDbResponse, utils.Tag, tagID)
		return
	}

	utils.MakeJsonResponse(w, http.StatusOK, tag)
}

func (h *handler) GetAllTagsByUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	userID := r.Header.Get("user_id")

	tags, err := h.service.Tag.GetAllTagsByUser(userID)
	if err != nil {
		utils.Abort(nil, w, http.StatusBadRequest, err, errors.ErrDbResponse, "", "")
		return
	}

	if len(tags) == 0 {
		utils.Abort(nil, w, http.StatusBadRequest, err, errors.ErrTagsListEmpty, "", "")
		return
	}

	utils.MakeJsonResponse(w, http.StatusOK, tags)
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
		utils.Abort(nil, w, http.StatusBadRequest, err, errors.ErrDbResponse, utils.Tag, tagID)
		return
	}

	err = h.service.Tag.UpdateTag(tag, tagID)
	if err != nil {
		utils.Abort(nil, w, http.StatusBadRequest, err, errors.ErrDbResponse, "", "")
		return
	}

	resp := make(map[string]string)
	resp["Обновлен тэг с id"] = tagID

	utils.MakeJsonResponse(w, http.StatusOK, resp)
}

func (h *handler) DeleteTag(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID := r.Header.Get("user_id")
	tagID := ps.ByName("id")

	_, err := h.service.Tag.GetTagByID(tagID, userID)
	if err != nil {
		utils.Abort(nil, w, http.StatusBadRequest, err, errors.ErrDbResponse, utils.Tag, tagID)
		return
	}

	id, err := h.service.Tag.DeleteTag(tagID, userID)
	if err != nil {
		utils.Abort(nil, w, http.StatusBadRequest, err, errors.ErrDbResponse, "", "")
		return
	}

	resp := make(map[string]int)
	resp["Удален тэг с id"] = id

	utils.MakeJsonResponse(w, http.StatusOK, resp)
}
