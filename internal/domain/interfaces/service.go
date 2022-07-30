package interfaces

import (
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
