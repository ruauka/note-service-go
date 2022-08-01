package services

import (
	"web/internal/adapters/storage"
	"web/internal/domain/enteties/dto"
	"web/internal/domain/enteties/model"
)

type UserAuthService interface {
	RegisterUser(user *model.User) (*model.User, error)
	GenerateToken(userName, password string) (string, error)
	ParseToken(token string) (string, error)
}

type UserService interface {
	GetAllUsers() ([]dto.UserResp, error)
	GetUserByID(id string) (*dto.UserResp, error)
	UpdateUser(newUser *dto.UserUpdate, userId string) error
	DeleteUser(id string) (int, error)
}

type NoteService interface {
	CreateNote(note *model.Note, userID string) (*model.Note, error)
	GetNoteByID(noteID, userID string) (*dto.NoteResp, error)
	GetAllNotesByUser(userID string) ([]dto.NotesResp, error)
	UpdateNote(newNote *dto.NoteUpdate, noteID string) error
	DeleteNote(noteID, userID string) (int, error)
}

type TagService interface {
	CreateTag(tag *model.Tag, userID string) (*model.Tag, error)
	GetTagByID(tagID, userID string) (*dto.TagResp, error)
	GetAllTagsByUser(userID string) ([]dto.TagsResp, error)
	UpdateTag(tag *dto.TagUpdate, tagID string) error
	DeleteTag(tagID, userID string) (int, error)
}

type Services struct {
	Auth UserAuthService
	User UserService
	Note NoteService
	Tag  TagService
}

func NewServices(db *storage.Storages) *Services {
	return &Services{
		Auth: NewAuthService(db.Auth),
		User: NewUserService(db.User),
		Note: NewNoteService(db.Note),
		Tag:  NewTagService(db.Tag),
	}
}
