package interfaces

import (
	"web/internal/domain/enteties/dto"
	"web/internal/domain/enteties/model"
)

type Storage interface {
	RegisterUser(user *model.User) (*model.User, error)
	GetUserForAuth(userName, password string) (*model.User, error)
	GetAllUsers() ([]dto.UserResp, error)
	GetUserByID(id string) (*dto.UserResp, error)
	UpdateUser(newUser *dto.UserUpdate, userId string) error
	DeleteUser(id string) (int, error)
}
