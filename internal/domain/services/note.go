package services

import (
	"web/internal/adapters/storage"
	"web/internal/domain/enteties/dto"
	"web/internal/domain/enteties/model"
)

// noteService note service struct.
type noteService struct {
	storage storage.NoteStorage
}

// NewNoteService note service func builder.
func NewNoteService(noteStorage storage.NoteStorage) NoteService {
	return &noteService{storage: noteStorage}
}

// CreateNote create user note.
func (n *noteService) CreateNote(note *model.Note, userID string) (*model.Note, error) {
	return n.storage.CreateNote(note, userID)
}

// GetNoteByID get note by ID.
func (n *noteService) GetNoteByID(noteID, userID string) (*dto.NoteResp, error) {
	return n.storage.GetNoteByID(noteID, userID)
}

// GetAllNotesByUser get all notes by user.
func (n *noteService) GetAllNotesByUser(userID string) ([]dto.NotesResp, error) {
	return n.storage.GetAllNotesByUser(userID)
}

// UpdateNote update note by ID.
func (n *noteService) UpdateNote(newNote *dto.NoteUpdate, noteID string) error {
	return n.storage.UpdateNote(newNote, noteID)
}

// DeleteNote delete note by ID.
func (n *noteService) DeleteNote(noteID, userID string) (int, error) {
	return n.storage.DeleteNote(noteID, userID)
}

// SetTags set tags to note.
func (n *noteService) SetTags(noteID string, tags map[string]string) (string, error) {
	return n.storage.SetTags(noteID, tags)
}

// RemoveTags remove tags from note.
func (n *noteService) RemoveTags(noteID string, tags map[string]string) (string, error) {
	return n.storage.RemoveTags(noteID, tags)
}

// GetAllNotesWithTags get all notes with tags by user.
func (n *noteService) GetAllNotesWithTags(userID string, notes []dto.NotesResp) ([]dto.NoteWithTagsResp, error) {
	return n.storage.GetAllNotesWithTags(userID, notes)
}

// GetNoteWithAllTags get note by id with all tags by user.
func (n *noteService) GetNoteWithAllTags(userID, noteID string, note *dto.NoteResp) (dto.NoteWithTagsResp, error) {
	return n.storage.GetNoteWithAllTags(userID, noteID, note)
}
