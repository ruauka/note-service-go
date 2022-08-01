package note

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

func (h *handler) GetAllNotesByUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	userID := r.Header.Get("user_id")

	notes, err := h.service.Note.GetAllNotesByUser(userID)
	if err != nil {
		utils.Abort(w, http.StatusBadRequest, err, errors.ErrDbResponse)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(notes)
}

func (h *handler) GetNoteByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID := r.Header.Get("user_id")

	note, err := h.service.Note.GetNoteByID(ps.ByName("id"), userID)
	if err != nil {
		utils.Abort(w, http.StatusBadRequest, err, errors.ErrDbResponse)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(note)
}

func (h *handler) CreateNote(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	note := &model.Note{}
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	userID := r.Header.Get("user_id")

	note, err := h.service.Note.CreateNote(note, userID)
	if err != nil {
		utils.Abort(w, http.StatusBadRequest, err, errors.ErrDbResponse)
		return
	}

	resp := make(map[string]string)
	resp[fmt.Sprintf("Создана заметка '%s' с id", note.Title)] = note.ID

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func (h *handler) UpdateNote(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	newNote := &dto.NoteUpdate{}
	if err := json.NewDecoder(r.Body).Decode(&newNote); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	userID := r.Header.Get("user_id")
	noteID := ps.ByName("id")

	_, err := h.service.Note.GetNoteByID(noteID, userID)
	if err != nil {
		utils.Abort(w, http.StatusBadRequest, err, errors.ErrNoteNotExists)
		return
	}

	err = h.service.Note.UpdateNote(newNote, noteID)
	if err != nil {
		utils.Abort(w, http.StatusBadRequest, err, errors.ErrDbResponse)
		return
	}

	resp := make(map[string]string)
	resp["Обновлена заметка с id"] = ps.ByName("id")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *handler) DeleteNote(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID := r.Header.Get("user_id")

	noteID, err := h.service.Note.DeleteNote(ps.ByName("id"), userID)
	if err != nil {
		utils.Abort(w, http.StatusBadRequest, err, errors.ErrDbResponse)
		return
	}

	resp := make(map[string]int)
	resp["Удалена заметка с id"] = noteID

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
