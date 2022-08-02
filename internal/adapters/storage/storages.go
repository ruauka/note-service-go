package storage

import (
	"github.com/jmoiron/sqlx"

	"web/internal/domain/enteties/dto"
	"web/internal/domain/enteties/model"
)

type UserAuthStorage interface {
	RegisterUser(user *model.User) (*model.User, error)
	GetUserForToken(userName, password string) (*model.User, error)
}

type UserStorage interface {
	GetAllUsers() ([]dto.UserResp, error)
	GetUserByID(id string) (*dto.UserResp, error)
	UpdateUser(newUser *dto.UserUpdate, userId string) error
	DeleteUser(id string) (int, error)
}

type NoteStorage interface {
	CreateNote(note *model.Note, userID string) (*model.Note, error)
	GetNoteByID(id string, userID string) (*dto.NoteResp, error)
	GetAllNotesByUser(userID string) ([]dto.NotesResp, error)
	UpdateNote(newNote *dto.NoteUpdate, noteID string) error
	DeleteNote(noteID, userID string) (int, error)
	SetTags(noteID string, tags []string) error
	RemoveTags(noteID string, tags []string) error
	GetAllNotesWithTags(userID string, notes []dto.NotesResp) ([]dto.NotesWithTagsResp, error)
}

type TagStorage interface {
	CreateTag(tag *model.Tag, userID string) (*model.Tag, error)
	GetTagByID(tagID, userID string) (*dto.TagResp, error)
	GetAllTagsByUser(userID string) ([]dto.TagsResp, error)
	UpdateTag(tag *dto.TagUpdate, tagID string) error
	DeleteTag(tagID, userID string) (int, error)
}

type Storages struct {
	Auth UserAuthStorage
	User UserStorage
	Note NoteStorage
	Tag  TagStorage
}

func NewStorages(pgDB *sqlx.DB) *Storages {
	return &Storages{
		Auth: NewAuthStorage(pgDB),
		User: NewUserStorage(pgDB),
		Note: NewNoteStorage(pgDB),
		Tag:  NewTagStorage(pgDB),
	}
}
