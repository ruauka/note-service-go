package storage

import (
	"github.com/jmoiron/sqlx"

	"web/internal/domain/enteties/model"
)

type noteStorage struct {
	db *sqlx.DB
}

func NewNoteStorage(pgDB *sqlx.DB) NoteStorage {
	return &noteStorage{db: pgDB}
}

func (n *noteStorage) CreateNote(note *model.Note) (*model.Note, error) {
	return &model.Note{ID: "11111"}, nil
}
