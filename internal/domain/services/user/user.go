package user

import (
	"web/internal/domain/enteties/dto"
	"web/internal/domain/enteties/model"
	"web/internal/domain/errors"
	"web/internal/domain/interfaces"
	"web/internal/domain/services/utils"
)

type userService struct {
	storage interfaces.Storage
	// logger
}

func NewUserService(db interfaces.Storage) interfaces.UserService {
	return &userService{storage: db}
}

func (u *userService) GetAllUsers() ([]dto.UserResponse, error) {
	users, err := u.storage.GetAllUsers()
	if len(users) == 0 {
		return nil, errors.ErrUsersListEmpty
	}
	return users, err
}

func (u *userService) GetUserByID(id string) (*model.User, error) {
	return u.storage.GetUserByID(id)
}

func (u *userService) CreateUser(user *model.User) (*model.User, error) {
	user.Password = utils.GeneratePasswordHash(user.Password)
	return u.storage.CreateUser(user)
}

func (u *userService) UpdateUser(newUser *dto.UserUpdate, userId string) error {
	return u.storage.UpdateUser(newUser, userId)
}

func (u *userService) DeleteUser(id string) (int, error) {
	return u.storage.DeleteUser(id)
}
