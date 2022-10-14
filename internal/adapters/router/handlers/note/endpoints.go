package note

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

// CreateNote create user note.
func (h *Handler) CreateNote(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	userID := r.Header.Get("user_id")
	ctx := r.Context()

	newNote := &model.Note{}
	if err := json.NewDecoder(r.Body).Decode(&newNote); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	err := validate.InputJSONValidate(newNote)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	note, err := h.service.Note.CreateNote(newNote, userID)
	if err != nil {
		functions.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrDBResponse, dictionary.Note, newNote.Title)
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	resp := make(map[string]string)
	resp[fmt.Sprintf("Created note '%s' with id", note.Title)] = note.ID

	functions.MakeJSONResponse(w, http.StatusCreated, resp)
}

// GetNoteByID get note by ID.
func (h *Handler) GetNoteByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID := r.Header.Get("user_id")
	noteID := ps.ByName("id")
	ctx := r.Context()

	note, err := h.service.Note.GetNoteByID(noteID, userID)
	if err != nil {
		functions.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrDBResponse, dictionary.Note, noteID)
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	functions.MakeJSONResponse(w, http.StatusOK, note)
}

// GetAllNotesByUser get all notes by user.
func (h *Handler) GetAllNotesByUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	userID := r.Header.Get("user_id")
	ctx := r.Context()

	notes, err := h.service.Note.GetAllNotesByUser(userID)
	if err != nil {
		functions.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrDBResponse, "", "")
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	if len(notes) == 0 {
		functions.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrNotesListEmpty, "", "")
		logger.LogFromContext(ctx).Error(errors.ErrNotesListEmpty.Error())
		return
	}

	functions.MakeJSONResponse(w, http.StatusOK, notes)
}

// UpdateNote update note by ID.
func (h *Handler) UpdateNote(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID := r.Header.Get("user_id")
	noteID := ps.ByName("id")
	ctx := r.Context()

	newNote := &dto.NoteUpdate{}
	if err := json.NewDecoder(r.Body).Decode(&newNote); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	_, err := h.service.Note.GetNoteByID(noteID, userID)
	if err != nil {
		functions.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrDBResponse, dictionary.Note, noteID)
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	err = h.service.Note.UpdateNote(newNote, noteID)
	if err != nil {
		functions.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrDBResponse, "", "")
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	resp := make(map[string]string)
	resp["Updated note with id"] = noteID

	functions.MakeJSONResponse(w, http.StatusOK, resp)
}

// DeleteNote delete note by ID.
func (h *Handler) DeleteNote(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID := r.Header.Get("user_id")
	noteID := ps.ByName("id")
	ctx := r.Context()

	_, err := h.service.Note.GetNoteByID(noteID, userID)
	if err != nil {
		functions.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrDBResponse, dictionary.Note, noteID)
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	id, err := h.service.Note.DeleteNote(noteID, userID)
	if err != nil {
		functions.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrDBResponse, "", "")
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	resp := make(map[string]int)
	resp["Deleted note with id"] = id

	functions.MakeJSONResponse(w, http.StatusOK, resp)
}

// SetTags set tags to note.
func (h *Handler) SetTags(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID := r.Header.Get("user_id")
	noteID := ps.ByName("id")
	tags := r.URL.Query()
	ctx := r.Context()

	note, err := h.service.Note.GetNoteByID(noteID, userID)
	if err != nil {
		functions.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrDBResponse, dictionary.Note, noteID)
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	tagsMap := make(map[string]string)

	for _, tagID := range tags["tag"] {
		tag, err := h.service.Tag.GetTagByID(tagID, userID)
		if err != nil {
			functions.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrDBResponse, dictionary.Tag, tagID)
			logger.LogFromContext(ctx).Error(err.Error())
			return
		}
		tagsMap[tagID] = tag.TagName
	}

	tagID, err := h.service.Note.SetTags(noteID, tagsMap)
	if err != nil {
		functions.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrDBResponse, dictionary.Tag, tagID)
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	resp := make(map[string]map[string]string)
	resp[fmt.Sprintf("To note '%s' set tags", note.Title)] = tagsMap

	functions.MakeJSONResponse(w, http.StatusOK, resp)
}

// RemoveTags remove tags from note.
func (h *Handler) RemoveTags(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID := r.Header.Get("user_id")
	noteID := ps.ByName("id")
	tags := r.URL.Query()
	ctx := r.Context()

	note, err := h.service.Note.GetNoteByID(noteID, userID)
	if err != nil {
		functions.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrDBResponse, dictionary.Note, noteID)
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	tagsMap := make(map[string]string)

	for _, tagID := range tags["tag"] {
		tag, err := h.service.Tag.GetTagByID(tagID, userID)
		if err != nil {
			functions.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrDBResponse, dictionary.Tag, tagID)
			logger.LogFromContext(ctx).Error(err.Error())
			return
		}
		tagsMap[tagID] = tag.TagName
	}

	tagID, err := h.service.Note.RemoveTags(noteID, tagsMap)
	if err != nil {
		functions.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrDBResponse, dictionary.Tag, tagID)
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	resp := make(map[string]map[string]string)
	resp[fmt.Sprintf("From note '%s' deleted tags", note.Title)] = tagsMap

	functions.MakeJSONResponse(w, http.StatusOK, resp)
}

// GetAllNotesWithTags get all notes with tags by user.
func (h *Handler) GetAllNotesWithTags(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	userID := r.Header.Get("user_id")
	ctx := r.Context()

	notes, err := h.service.Note.GetAllNotesByUser(userID)
	if err != nil {
		functions.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrDBResponse, "", "")
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	if len(notes) == 0 {
		functions.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrNotesListEmpty, "", "")
		logger.LogFromContext(ctx).Error(errors.ErrNotesListEmpty.Error())
		return
	}

	notesResp, err := h.service.Note.GetAllNotesWithTags(userID, notes)
	if err != nil {
		functions.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrDBResponse, "", "")
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	if len(notesResp) == 0 {
		functions.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrNotesListWithTagsEmpty, "", "")
		logger.LogFromContext(ctx).Error(errors.ErrNotesListWithTagsEmpty.Error())
		return
	}

	functions.MakeJSONResponse(w, http.StatusOK, notesResp)
}

// GetNoteWithAllTags get note by id with all tags by user.
func (h *Handler) GetNoteWithAllTags(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID := r.Header.Get("user_id")
	noteID := ps.ByName("id")
	ctx := r.Context()

	note, err := h.service.Note.GetNoteByID(noteID, userID)
	if err != nil {
		functions.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrDBResponse, dictionary.Note, noteID)
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	noteResp, err := h.service.Note.GetNoteWithAllTags(userID, noteID, note)
	if err != nil {
		functions.Abort(ctx, w, http.StatusBadRequest, nil, err, "", "")
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	functions.MakeJSONResponse(w, http.StatusOK, noteResp)
}
