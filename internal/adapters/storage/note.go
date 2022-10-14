package storage

import (
	"fmt"
	"log"
	"strings"

	"github.com/jmoiron/sqlx"

	"web/internal/domain/entities/dto"
	"web/internal/domain/entities/model"
	"web/internal/utils/dictionary"
)

// noteStorage note storage struct.
type noteStorage struct {
	db *sqlx.DB
}

// NewNoteStorage note storage func builder.
func NewNoteStorage(pgDB *sqlx.DB) NoteStorage {
	return &noteStorage{db: pgDB}
}

// CreateNote create note in DB.
func (n *noteStorage) CreateNote(note *model.Note, userID string) (*model.Note, error) {
	query := fmt.Sprintf("INSERT INTO %s (title, info, user_id) VALUES ($1, $2, $3) RETURNING id", dictionary.NotesTable)
	if err := n.db.QueryRow(query, note.Title, note.Info, userID).Scan(&note.ID); err != nil {
		return nil, err
	}

	return note, nil
}

// GetNoteByID get note by id from DB.
func (n *noteStorage) GetNoteByID(id string, userID string) (*dto.NoteResp, error) {
	var note dto.NoteResp

	query := fmt.Sprintf("SELECT title, info FROM %s WHERE id=$1 AND user_id=$2", dictionary.NotesTable)
	if err := n.db.Get(&note, query, id, userID); err != nil {
		return nil, err
	}

	return &note, nil
}

// GetAllNotesByUser get all notes by user from DB.
func (n *noteStorage) GetAllNotesByUser(userID string) ([]dto.NotesResp, error) {
	var notes []dto.NotesResp

	query := fmt.Sprintf("SELECT id, title, info FROM %s WHERE user_id=$1", dictionary.NotesTable)
	if err := n.db.Select(&notes, query, userID); err != nil {
		return nil, err
	}

	return notes, nil
}

// UpdateNote update note by id in DB.
func (n *noteStorage) UpdateNote(newNote *dto.NoteUpdate, noteID string) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 1

	if newNote.Info != nil {
		setValues = append(setValues, fmt.Sprintf("info=$%d", argID))
		args = append(args, *newNote.Info)
		argID++
	}

	if newNote.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argID))
		args = append(args, *newNote.Title)
		argID++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", dictionary.NotesTable, setQuery, argID)
	args = append(args, noteID)

	_, err := n.db.Exec(query, args...)

	return err
}

// DeleteNote delete note by id from DB.
func (n *noteStorage) DeleteNote(noteID, userID string) (int, error) {
	var id int

	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1 AND user_id=$2 RETURNING id", dictionary.NotesTable)
	if err := n.db.QueryRow(query, noteID, userID).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

// SetTags - set tags to note by ID.
func (n *noteStorage) SetTags(noteID string, tags map[string]string) (string, error) {
	for tagID, tagName := range tags {
		query := fmt.Sprintf("INSERT INTO %s (note_id, tag_id) VALUES ($1, $2)", dictionary.NotesTagsTable)
		if res := n.db.QueryRow(query, noteID, tagID); res.Err() != nil {
			return tagName, res.Err()
		}
	}

	return "", nil
}

// RemoveTags - remove tags from note by ID.
func (n *noteStorage) RemoveTags(noteID string, tags map[string]string) (string, error) {
	for tagID, tagName := range tags {
		query := fmt.Sprintf("DELETE FROM %s WHERE note_id=$1 AND tag_id=$2", dictionary.NotesTagsTable)
		if res := n.db.QueryRow(query, noteID, tagID); res.Err() != nil {
			return tagName, res.Err()
		}
	}

	return "", nil
}

// GetAllNotesWithTags get all notes with tags by user.
func (n *noteStorage) GetAllNotesWithTags(userID string, notes []dto.NotesResp) ([]dto.NoteWithTagsResp, error) {
	tempNotes := make([]dto.NoteWithTagsResp, len(notes))
	resultNotes := make([]dto.NoteWithTagsResp, 0)

	for index, note := range notes {
		query := fmt.Sprintf("SELECT notes.id, notes.title, notes.info, tags.id, tags.tagname"+
			" FROM %s JOIN notes_tags"+
			" ON notes.id = notes_tags.note_id"+
			" JOIN tags"+
			" ON notes_tags.tag_id = tags.id"+
			" AND tags.user_id = $1 AND notes.id = $2", dictionary.NotesTable)

		row, err := n.db.Query(query, userID, note.ID) //nolint:rowserrcheck
		if err != nil {
			return nil, err
		}

		for row.Next() {
			var note dto.NoteWithTags

			if err := row.Scan(&note.ID, &note.Title, &note.Info, &note.TagsResp.ID, &note.TagsResp.TagName); err != nil {
				log.Println(err)
			}

			if note.ID != "" {
				tempNotes[index].ID = note.ID
				tempNotes[index].Title = note.Title
				tempNotes[index].Info = note.Info
				tempNotes[index].TagsResp = append(tempNotes[index].TagsResp, note.TagsResp)
			}
		}
	}

	for _, note := range tempNotes {
		if note.ID != "" {
			resultNotes = append(resultNotes, note)
		}
	}

	return resultNotes, nil
}

// GetNoteWithAllTags get note by id with all tags by user.
func (n *noteStorage) GetNoteWithAllTags(userID, noteID string, note *dto.NoteResp) (dto.NoteWithTagsResp, error) {
	var resultNote dto.NoteWithTagsResp

	query := fmt.Sprintf("SELECT notes.id, notes.title, notes.info, tags.id, tags.tagname"+
		" FROM %s JOIN notes_tags"+
		" ON notes.id = notes_tags.note_id"+
		" JOIN tags"+
		" ON notes_tags.tag_id = tags.id"+
		" AND tags.user_id = $1 AND notes.id = $2", dictionary.NotesTable)

	row, err := n.db.Query(query, userID, noteID) //nolint:rowserrcheck
	if err != nil {
		return dto.NoteWithTagsResp{}, err
	}

	for row.Next() {
		var note dto.NoteWithTags

		if err := row.Scan(&note.ID, &note.Title, &note.Info, &note.TagsResp.ID, &note.TagsResp.TagName); err != nil {
			log.Println(err)
		}

		resultNote.ID = note.ID
		resultNote.Title = note.Title
		resultNote.Info = note.Info
		resultNote.TagsResp = append(resultNote.TagsResp, note.TagsResp)
	}

	if resultNote.ID == "" {
		return dto.NoteWithTagsResp{}, fmt.Errorf(fmt.Sprintf("Note with id '%s' has no tags", noteID))
	}

	return resultNote, nil
}
