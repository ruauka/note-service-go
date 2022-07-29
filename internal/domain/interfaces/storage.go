package interfaces

import (
	"web/internal/domain/enteties/dto"
	"web/internal/domain/enteties/model"
)

type Storage interface {
	GetAllUsers() ([]dto.UserResponse, error)
	GetUserByID(id string) (*model.User, error)
	CreateUser(user *model.User) (*model.User, error)
	UpdateUser(newUser *dto.UserUpdate, userId string) error
	DeleteUser(id string) (int, error)
}
