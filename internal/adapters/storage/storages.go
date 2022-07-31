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
	GetAllNotesByUser(userID string) ([]dto.NoteResp, error)

	DeleteNote(noteID, userID string) (int, error)
}

type Storages struct {
	Auth UserAuthStorage
	User UserStorage
	Note NoteStorage
}

func NewStorages(pgDB *sqlx.DB) *Storages {
	return &Storages{
		Auth: NewAuthStorage(pgDB),
		User: NewUserStorage(pgDB),
		Note: NewNoteStorage(pgDB),
	}
}
