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

func NewNoteService(noteStorage storage.NoteStorage) NoteService {
	return &noteService{storage: noteStorage}
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

func (n *noteService) SetTags(noteID string, tags []string) error {
	return n.storage.SetTags(noteID, tags)
}

func (n *noteService) RemoveTags(noteID string, tags []string) error {
	return n.storage.RemoveTags(noteID, tags)
}

func (n *noteService) GetAllNotesWithTags(userID string, notes []dto.NotesResp) ([]dto.NotesWithTagsResp, error) {
	return n.storage.GetAllNotesWithTags(userID, notes)

}
