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
	"web/pkg/logger"
)

func (h *handler) CreateTag(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	userID := r.Header.Get("user_id")
	ctx := r.Context()

	newTag := &model.Tag{}
	if err := json.NewDecoder(r.Body).Decode(&newTag); err != nil {
		http.Error(w, err.Error(), 400)
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}
	// Валидация объекта структуры Tag //
	err := validate.InputJsonValidate(newTag)
	if err != nil {
		http.Error(w, err.Error(), 400)
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	tag, err := h.service.Tag.CreateTag(newTag, userID)
	if err != nil {
		utils.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrDbResponse, utils.Tag, newTag.TagName)
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	resp := make(map[string]string)
	resp[fmt.Sprintf("Created tag '%s' with id", tag.TagName)] = tag.ID

	utils.MakeJsonResponse(w, http.StatusCreated, resp)
}

func (h *handler) GetTagByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID := r.Header.Get("user_id")
	tagID := ps.ByName("id")
	ctx := r.Context()

	tag, err := h.service.Tag.GetTagByID(tagID, userID)
	if err != nil {
		utils.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrDbResponse, utils.Tag, tagID)
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	utils.MakeJsonResponse(w, http.StatusOK, tag)
}

func (h *handler) GetAllTagsByUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	userID := r.Header.Get("user_id")
	ctx := r.Context()

	tags, err := h.service.Tag.GetAllTagsByUser(userID)
	if err != nil {
		utils.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrDbResponse, "", "")
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	if len(tags) == 0 {
		utils.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrTagsListEmpty, "", "")
		logger.LogFromContext(ctx).Error(errors.ErrTagsListEmpty.Error())
		return
	}

	utils.MakeJsonResponse(w, http.StatusOK, tags)
}

func (h *handler) UpdateTag(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID := r.Header.Get("user_id")
	tagID := ps.ByName("id")
	ctx := r.Context()

	tag := &dto.TagUpdate{}
	if err := json.NewDecoder(r.Body).Decode(&tag); err != nil {
		http.Error(w, err.Error(), 400)
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	_, err := h.service.Tag.GetTagByID(tagID, userID)
	if err != nil {
		utils.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrDbResponse, utils.Tag, tagID)
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	err = h.service.Tag.UpdateTag(tag, tagID)
	if err != nil {
		utils.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrDbResponse, "", "")
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	resp := make(map[string]string)
	resp["Updated tag with id"] = tagID

	utils.MakeJsonResponse(w, http.StatusOK, resp)
}

func (h *handler) DeleteTag(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID := r.Header.Get("user_id")
	tagID := ps.ByName("id")
	ctx := r.Context()

	_, err := h.service.Tag.GetTagByID(tagID, userID)
	if err != nil {
		utils.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrDbResponse, utils.Tag, tagID)
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	id, err := h.service.Tag.DeleteTag(tagID, userID)
	if err != nil {
		utils.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrDbResponse, "", "")
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	resp := make(map[string]int)
	resp["Delete tag with id"] = id

	utils.MakeJsonResponse(w, http.StatusOK, resp)
}
