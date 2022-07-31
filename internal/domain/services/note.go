package services

import (
	"web/internal/adapters/storage"
	"web/internal/domain/enteties/model"
)

type noteService struct {
	storage storage.NoteStorage
	// logger
}

func NewNoteService(db storage.NoteStorage) NoteService {
	return &noteService{storage: db}
}

func (n *noteService) CreateNote(note *model.Note) (*model.Note, error) {
	return n.storage.CreateNote(note)
}
