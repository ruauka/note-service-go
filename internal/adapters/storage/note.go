package storage

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"web/internal/domain/enteties/model"
	"web/internal/utils"
)

type noteStorage struct {
	db *sqlx.DB
}

func NewNoteStorage(pgDB *sqlx.DB) NoteStorage {
	return &noteStorage{db: pgDB}
}

func (n *noteStorage) CreateNote(note *model.Note, userID string) (*model.Note, error) {
	query := fmt.Sprintf("INSERT INTO %s (note, user_id) VALUES ($1, $2) RETURNING id", utils.NotesTable)
	if err := n.db.QueryRow(query, note.Note, userID).Scan(&note.ID); err != nil {
		return nil, err
	}

	return note, nil
}
