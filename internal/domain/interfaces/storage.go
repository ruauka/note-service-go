package interfaces

import (
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
