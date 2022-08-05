package note

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

func (h *handler) GetAllNotesByUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	userID := r.Header.Get("user_id")
	ctx := r.Context()

	notes, err := h.service.Note.GetAllNotesByUser(userID)
	if err != nil {
		utils.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrDbResponse, "", "")
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	if len(notes) == 0 {
		utils.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrNotesListEmpty, "", "")
		logger.LogFromContext(ctx).Error(errors.ErrNotesListEmpty.Error())
		return
	}

	utils.MakeJsonResponse(w, http.StatusOK, notes)
}

func (h *handler) GetNoteByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID := r.Header.Get("user_id")
	noteID := ps.ByName("id")
	ctx := r.Context()

	note, err := h.service.Note.GetNoteByID(noteID, userID)
	if err != nil {
		utils.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrDbResponse, utils.Note, noteID)
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	utils.MakeJsonResponse(w, http.StatusOK, note)
}

func (h *handler) CreateNote(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	userID := r.Header.Get("user_id")
	ctx := r.Context()

	newNote := &model.Note{}
	if err := json.NewDecoder(r.Body).Decode(&newNote); err != nil {
		http.Error(w, err.Error(), 400)
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}
	// Валидация объекта структуры Note //
	err := validate.InputJsonValidate(newNote)
	if err != nil {
		http.Error(w, err.Error(), 400)
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	note, err := h.service.Note.CreateNote(newNote, userID)
	if err != nil {
		utils.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrDbResponse, utils.Note, newNote.Title)
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	resp := make(map[string]string)
	resp[fmt.Sprintf("Created note '%s' with id", note.Title)] = note.ID

	utils.MakeJsonResponse(w, http.StatusCreated, resp)
}

func (h *handler) UpdateNote(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID := r.Header.Get("user_id")
	noteID := ps.ByName("id")
	ctx := r.Context()

	newNote := &dto.NoteUpdate{}
	if err := json.NewDecoder(r.Body).Decode(&newNote); err != nil {
		http.Error(w, err.Error(), 400)
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	_, err := h.service.Note.GetNoteByID(noteID, userID)
	if err != nil {
		utils.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrDbResponse, utils.Note, noteID)
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	err = h.service.Note.UpdateNote(newNote, noteID)
	if err != nil {
		utils.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrDbResponse, "", "")
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	resp := make(map[string]string)
	resp["Updated note with id"] = noteID

	utils.MakeJsonResponse(w, http.StatusOK, resp)
}

func (h *handler) DeleteNote(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID := r.Header.Get("user_id")
	noteID := ps.ByName("id")
	ctx := r.Context()

	_, err := h.service.Note.GetNoteByID(noteID, userID)
	if err != nil {
		utils.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrDbResponse, utils.Note, noteID)
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	id, err := h.service.Note.DeleteNote(noteID, userID)
	if err != nil {
		utils.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrDbResponse, "", "")
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	resp := make(map[string]int)
	resp["Deleted note with id"] = id

	utils.MakeJsonResponse(w, http.StatusOK, resp)
}

func (h *handler) SetTags(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID := r.Header.Get("user_id")
	noteID := ps.ByName("id")
	tags := r.URL.Query()
	ctx := r.Context()

	note, err := h.service.Note.GetNoteByID(noteID, userID)
	if err != nil {
		utils.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrDbResponse, utils.Note, noteID)
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	tagsIDs := make([]string, 0, len(tags))

	for _, tagID := range tags {
		_, err := h.service.Tag.GetTagByID(tagID[0], userID)
		if err != nil {
			utils.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrDbResponse, utils.Tag, tagID[0])
			logger.LogFromContext(ctx).Error(err.Error())
			return
		}
		tagsIDs = append(tagsIDs, tagID[0])
	}

	if err := h.service.Note.SetTags(noteID, tagsIDs); err != nil {
		utils.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrDbResponse, "", "")
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	tagResp := make(map[string]string)
	for tag, id := range tags {
		tagResp[tag] = id[0]
	}

	resp := make(map[string]map[string]string)
	resp[fmt.Sprintf("К заметке '%s' добавлены тэги", note.Title)] = tagResp

	utils.MakeJsonResponse(w, http.StatusOK, resp)
}

func (h *handler) RemoveTags(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID := r.Header.Get("user_id")
	noteID := ps.ByName("id")
	tags := r.URL.Query()
	ctx := r.Context()

	note, err := h.service.Note.GetNoteByID(noteID, userID)
	if err != nil {
		utils.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrDbResponse, utils.Note, noteID)
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	tagsIDs := make([]string, 0, len(tags))

	for _, tagID := range tags {
		_, err := h.service.Tag.GetTagByID(tagID[0], userID)
		if err != nil {
			utils.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrDbResponse, utils.Tag, tagID[0])
			logger.LogFromContext(ctx).Error(err.Error())
			return
		}
		tagsIDs = append(tagsIDs, tagID[0])
	}

	if err := h.service.Note.RemoveTags(noteID, tagsIDs); err != nil {
		utils.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrDbResponse, "", "")
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	tagResp := make(map[string]string)
	for tag, id := range tags {
		tagResp[tag] = id[0]
	}

	resp := make(map[string]map[string]string)
	resp[fmt.Sprintf("У заметки '%s' удалены тэги", note.Title)] = tagResp

	utils.MakeJsonResponse(w, http.StatusOK, resp)
}

func (h *handler) GetAllNotesWithTags(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	userID := r.Header.Get("user_id")
	ctx := r.Context()

	notes, err := h.service.Note.GetAllNotesByUser(userID)
	if err != nil {
		utils.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrDbResponse, "", "")
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	if len(notes) == 0 {
		utils.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrNotesListEmpty, "", "")
		logger.LogFromContext(ctx).Error(errors.ErrNotesListEmpty.Error())
		return
	}

	notesResp, err := h.service.Note.GetAllNotesWithTags(userID, notes)
	if err != nil {
		utils.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrDbResponse, "", "")
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	if len(notesResp) == 0 {
		utils.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrNotesListWithTagsEmpty, "", "")
		logger.LogFromContext(ctx).Error(errors.ErrNotesListWithTagsEmpty.Error())
		return
	}

	utils.MakeJsonResponse(w, http.StatusOK, notesResp)
}

func (h *handler) GetNoteWithAllTags(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID := r.Header.Get("user_id")
	noteID := ps.ByName("id")
	ctx := r.Context()

	note, err := h.service.Note.GetNoteByID(noteID, userID)
	if err != nil {
		utils.Abort(ctx, w, http.StatusBadRequest, err, errors.ErrDbResponse, utils.Note, noteID)
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	noteResp, err := h.service.Note.GetNoteWithAllTags(userID, noteID, note)
	if err != nil {
		utils.Abort(ctx, w, http.StatusBadRequest, nil, err, "", "")
		logger.LogFromContext(ctx).Error(err.Error())
		return
	}

	utils.MakeJsonResponse(w, http.StatusOK, noteResp)
}
