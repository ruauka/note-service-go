// Package services Package services
package services

import (
	"web/internal/adapters/storage"
	"web/internal/domain/enteties/dto"
	"web/internal/domain/enteties/model"
)

//go:generate mockgen -source=services.go -destination=mocks/mock.go

// UserAuthService Auth interface.
type UserAuthService interface {
	RegisterUser(user *model.User) (*model.User, error)
	GenerateToken(userName, password string) (string, error)
	ParseToken(token string) (string, error)
}

// UserService User interface.
type UserService interface {
	GetUserByID(id string) (*dto.UserResp, error)
	GetAllUsers() ([]dto.UserResp, error)
	UpdateUser(newUser *dto.UserUpdate, userID string) error
	DeleteUser(id string) (int, error)
}

// NoteService Note interface.
type NoteService interface {
	CreateNote(note *model.Note, userID string) (*model.Note, error)
	GetNoteByID(noteID, userID string) (*dto.NoteResp, error)
	GetAllNotesByUser(userID string) ([]dto.NotesResp, error)
	UpdateNote(newNote *dto.NoteUpdate, noteID string) error
	DeleteNote(noteID, userID string) (int, error)
	SetTags(noteID string, tags map[string]string) (string, error)
	RemoveTags(noteID string, tags map[string]string) (string, error)
	GetAllNotesWithTags(userID string, notes []dto.NotesResp) ([]dto.NoteWithTagsResp, error)
	GetNoteWithAllTags(userID, NoteID string, note *dto.NoteResp) (dto.NoteWithTagsResp, error)
}

// TagService Tag interface.
type TagService interface {
	CreateTag(tag *model.Tag, userID string) (*model.Tag, error)
	GetTagByID(tagID, userID string) (*dto.TagResp, error)
	GetAllTagsByUser(userID string) ([]dto.TagsResp, error)
	UpdateTag(tag *dto.TagUpdate, tagID string) error
	DeleteTag(tagID, userID string) (int, error)
}

// Services struct of services interfaces.
type Services struct {
	Auth UserAuthService
	User UserService
	Note NoteService
	Tag  TagService
}

// NewServices services func builder.
func NewServices(storages *storage.Storages) *Services {
	return &Services{
		Auth: NewAuthService(storages.Auth),
		User: NewUserService(storages.User),
		Note: NewNoteService(storages.Note),
		Tag:  NewTagService(storages.Tag),
	}
}
