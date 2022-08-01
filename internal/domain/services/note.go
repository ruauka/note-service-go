package services

import (
	"web/internal/adapters/storage"
	"web/internal/domain/enteties/dto"
	"web/internal/domain/enteties/model"
)

type noteService struct {
	storage storage.NoteStorage
	// logger
}

func NewNoteService(db storage.NoteStorage) NoteService {
	return &noteService{storage: db}
}

func (n *noteService) CreateNote(note *model.Note, userID string) (*model.Note, error) {
	return n.storage.CreateNote(note, userID)
}

func (n *noteService) GetNoteByID(noteID, userID string) (*dto.NoteResp, error) {
	return n.storage.GetNoteByID(noteID, userID)
}

func (n *noteService) GetAllNotesByUser(userID string) ([]dto.NotesResp, error) {
	return n.storage.GetAllNotesByUser(userID)
}

func (n *noteService) UpdateNote(newNote *dto.NoteUpdate, noteID string) error {
	return n.storage.UpdateNote(newNote, noteID)
}

func (n *noteService) DeleteNote(noteID, userID string) (int, error) {
	return n.storage.DeleteNote(noteID, userID)
}