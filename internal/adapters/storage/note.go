package storage

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"web/internal/domain/enteties/dto"
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
	query := fmt.Sprintf("INSERT INTO %s (title, info, user_id) VALUES ($1, $2, $3) RETURNING id", utils.NotesTable)
	if err := n.db.QueryRow(query, note.Title, note.Info, userID).Scan(&note.ID); err != nil {
		return nil, err
	}

	return note, nil
}

func (n *noteStorage) GetNoteByID(id string, userID string) (*dto.NoteResp, error) {
	var note dto.NoteResp

	query := fmt.Sprintf("SELECT title, info FROM %s WHERE id=$1 AND user_id=$2", utils.NotesTable)
	if err := n.db.Get(&note, query, id, userID); err != nil {
		return nil, err
	}

	return &note, nil
}

func (n *noteStorage) GetAllNotesByUser(userID string) ([]dto.NoteResp, error) {
	var notes []dto.NoteResp

	query := fmt.Sprintf("SELECT title, info FROM %s WHERE user_id=$1", utils.NotesTable)
	if err := n.db.Select(&notes, query, userID); err != nil {
		return nil, err
	}

	return notes, nil
}

func (n *noteStorage) DeleteNote(noteID, userID string) (int, error) {
	var id int

	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1 AND user_id=$2 RETURNING id", utils.NotesTable)
	if err := n.db.QueryRow(query, noteID, userID).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
