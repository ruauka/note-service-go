package storage

import (
	"fmt"
	"log"
	"strings"

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

func (n *noteStorage) GetAllNotesByUser(userID string) ([]dto.NotesResp, error) {
	var notes []dto.NotesResp

	query := fmt.Sprintf("SELECT id, title, info FROM %s WHERE user_id=$1", utils.NotesTable)
	if err := n.db.Select(&notes, query, userID); err != nil {
		return nil, err
	}

	return notes, nil
}

func (n *noteStorage) UpdateNote(newNote *dto.NoteUpdate, noteID string) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if newNote.Info != nil {
		setValues = append(setValues, fmt.Sprintf("info=$%d", argId))
		args = append(args, *newNote.Info)
		argId++
	}

	if newNote.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *newNote.Title)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", utils.NotesTable, setQuery, argId)
	args = append(args, noteID)

	_, err := n.db.Exec(query, args...)

	return err
}

func (n *noteStorage) DeleteNote(noteID, userID string) (int, error) {
	var id int

	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1 AND user_id=$2 RETURNING id", utils.NotesTable)
	if err := n.db.QueryRow(query, noteID, userID).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (n *noteStorage) SetTags(noteID string, tags []string) error {
	for _, tagID := range tags {
		query := fmt.Sprintf("INSERT INTO %s (note_id, tag_id) VALUES ($1, $2)", utils.NotesTagsTable)
		if res := n.db.QueryRow(query, noteID, tagID); res.Err() != nil {
			return res.Err()
		}
	}

	return nil
}

func (n *noteStorage) RemoveTags(noteID string, tags []string) error {
	for _, tagID := range tags {
		query := fmt.Sprintf("DELETE FROM %s WHERE note_id=$1 AND tag_id=$2", utils.NotesTagsTable)
		if res := n.db.QueryRow(query, noteID, tagID); res.Err() != nil {
			return res.Err()
		}
	}

	return nil
}

func (n *noteStorage) GetAllNotesWithTags(userID string, notes []dto.NotesResp) ([]dto.NotesWithTagsResp, error) {
	resultNotes := make([]dto.NotesWithTagsResp, len(notes), len(notes))

	for index, note := range notes {
		query := fmt.Sprintf("SELECT notes.id, notes.title, notes.info, tags.id, tags.tagname"+
			" FROM %s JOIN notes_tags"+
			" ON notes.id = notes_tags.note_id"+
			" JOIN tags"+
			" ON notes_tags.tag_id = tags.id"+
			" AND tags.user_id = $1 AND notes.id = $2", utils.NotesTable)

		row, err := n.db.Query(query, userID, note.ID)
		if err != nil {
			fmt.Println(err)
		}

		for row.Next() {
			var note dto.NotesWithTags

			if err := row.Scan(&note.ID, &note.Title, &note.Info, &note.TagsResp.ID, &note.TagsResp.TagName); err != nil {
				log.Println(err)
			}

			resultNotes[index].ID = note.ID
			resultNotes[index].Title = note.Title
			resultNotes[index].Info = note.Info
			resultNotes[index].TagsResp = append(resultNotes[index].TagsResp, note.TagsResp)
		}
	}

	return resultNotes, nil
}
