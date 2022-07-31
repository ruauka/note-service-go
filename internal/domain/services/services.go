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
	CreateNote(note *model.Note) (*model.Note, error)
}

type Services struct {
	Auth UserAuthService
	User UserService
	Note NoteService
}

func NewServices(db *storage.Storages) *Services {
	return &Services{
		Auth: NewAuthService(db.Auth),
		User: NewUserService(db.User),
		Note: NewNoteService(db.Note),
	}
}
