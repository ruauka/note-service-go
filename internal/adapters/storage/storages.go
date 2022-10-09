// Package storage Package storage
package storage

import (
	"github.com/jmoiron/sqlx"

	"web/internal/domain/entities/dto"
	"web/internal/domain/entities/model"
)

// UserAuthStorage Auth interface.
type UserAuthStorage interface {
	RegisterUser(user *model.User) (*model.User, error)
	GetUserForToken(userName, password string) (*model.User, error)
}

// UserStorage User interface.
type UserStorage interface {
	GetUserByID(id string) (*dto.UserResp, error)
	GetAllUsers() ([]dto.UserResp, error)
	UpdateUser(newUser *dto.UserUpdate, userID string) error
	DeleteUser(id string) (int, error)
}

// NoteStorage Note interface.
type NoteStorage interface {
	CreateNote(note *model.Note, userID string) (*model.Note, error)
	GetNoteByID(id string, userID string) (*dto.NoteResp, error)
	GetAllNotesByUser(userID string) ([]dto.NotesResp, error)
	UpdateNote(newNote *dto.NoteUpdate, noteID string) error
	DeleteNote(noteID, userID string) (int, error)
	SetTags(noteID string, tags map[string]string) (string, error)
	RemoveTags(noteID string, tags map[string]string) (string, error)
	GetAllNotesWithTags(userID string, notes []dto.NotesResp) ([]dto.NoteWithTagsResp, error)
	GetNoteWithAllTags(userID, noteID string, note *dto.NoteResp) (dto.NoteWithTagsResp, error)
}

// TagStorage Tag interface.
type TagStorage interface {
	CreateTag(tag *model.Tag, userID string) (*model.Tag, error)
	GetTagByID(tagID, userID string) (*dto.TagResp, error)
	GetAllTagsByUser(userID string) ([]dto.TagsResp, error)
	UpdateTag(tag *dto.TagUpdate, tagID string) error
	DeleteTag(tagID, userID string) (int, error)
}

// Storages struct of storages interfaces.
type Storages struct {
	Auth UserAuthStorage
	User UserStorage
	Note NoteStorage
	Tag  TagStorage
}

// NewStorages storages func builder.
func NewStorages(pgDB *sqlx.DB) *Storages {
	return &Storages{
		Auth: NewAuthStorage(pgDB),
		User: NewUserStorage(pgDB),
		Note: NewNoteStorage(pgDB),
		Tag:  NewTagStorage(pgDB),
	}
}
