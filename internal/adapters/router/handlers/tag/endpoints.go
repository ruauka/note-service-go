// Package tag Package tag
package tag

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

// CreateTag create tag.
func (h *Handler) CreateTag(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	userID := r.Header.Get("user_id")
	ctx := r.Context()

	newTag := &model.Tag{}
	if err := json.NewDecoder(r.Body).Decode(&newTag); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}
	// Валидация объекта структуры Tag //
	err := validate.InputJSONValidate(newTag)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	tag, err := h.service.Tag.CreateTag(newTag, userID)
	if err != nil {
		functions.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrDBResponse, dictionary.Tag, newTag.TagName)
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	resp := make(map[string]string)
	resp[fmt.Sprintf("Created tag '%s' with id", tag.TagName)] = tag.ID

	functions.MakeJSONResponse(w, http.StatusCreated, resp)
}

// GetTagByID get tag by ID.
func (h *Handler) GetTagByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID := r.Header.Get("user_id")
	tagID := ps.ByName("id")
	ctx := r.Context()

	tag, err := h.service.Tag.GetTagByID(tagID, userID)
	if err != nil {
		functions.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrDBResponse, dictionary.Tag, tagID)
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	functions.MakeJSONResponse(w, http.StatusOK, tag)
}

// GetAllTagsByUser get tag by user.
func (h *Handler) GetAllTagsByUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	userID := r.Header.Get("user_id")
	ctx := r.Context()

	tags, err := h.service.Tag.GetAllTagsByUser(userID)
	if err != nil {
		functions.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrDBResponse, "", "")
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	if len(tags) == 0 {
		functions.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrTagsListEmpty, "", "")
		logger.LogFromContext(ctx).Error(errors.ErrTagsListEmpty.Error())
		return
	}

	functions.MakeJSONResponse(w, http.StatusOK, tags)
}

// UpdateTag update tag by ID.
func (h *Handler) UpdateTag(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID := r.Header.Get("user_id")
	tagID := ps.ByName("id")
	ctx := r.Context()

	tag := &dto.TagUpdate{}
	if err := json.NewDecoder(r.Body).Decode(&tag); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	_, err := h.service.Tag.GetTagByID(tagID, userID)
	if err != nil {
		functions.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrDBResponse, dictionary.Tag, tagID)
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	err = h.service.Tag.UpdateTag(tag, tagID)
	if err != nil {
		functions.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrDBResponse, "", "")
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	resp := make(map[string]string)
	resp["Updated tag with id"] = tagID

	functions.MakeJSONResponse(w, http.StatusOK, resp)
}

// DeleteTag delete tag by ID.
func (h *Handler) DeleteTag(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID := r.Header.Get("user_id")
	tagID := ps.ByName("id")
	ctx := r.Context()

	_, err := h.service.Tag.GetTagByID(tagID, userID)
	if err != nil {
		functions.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrDBResponse, dictionary.Tag, tagID)
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	id, err := h.service.Tag.DeleteTag(tagID, userID)
	if err != nil {
		functions.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrDBResponse, "", "")
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	resp := make(map[string]int)
	resp["Deleted tag with id"] = id

	functions.MakeJSONResponse(w, http.StatusOK, resp)
}
